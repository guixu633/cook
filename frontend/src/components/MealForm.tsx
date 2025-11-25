import React, { useState } from 'react';
import { Form, Input, Button, DatePicker, message, Upload, Modal } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import type { UploadFile, UploadProps } from 'antd';
import type { CreateMealRequest } from '../types/meal';
import { uploadImage } from '../api/meal';
import dayjs from 'dayjs';

interface MealFormProps {
  onSubmit: (values: CreateMealRequest) => Promise<void>;
}

export const MealForm: React.FC<MealFormProps> = ({ onSubmit }) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  
  // State for image uploads
  const [fileList, setFileList] = useState<UploadFile[]>([]);
  const [previewOpen, setPreviewOpen] = useState(false);
  const [previewImage, setPreviewImage] = useState('');
  const [previewTitle, setPreviewTitle] = useState('');

  const handleCancel = () => setPreviewOpen(false);

  const handlePreview = async (file: UploadFile) => {
    if (!file.url && !file.preview) {
      file.preview = await getBase64(file.originFileObj as File);
    }

    setPreviewImage(file.url || (file.preview as string));
    setPreviewOpen(true);
    setPreviewTitle(file.name || file.url!.substring(file.url!.lastIndexOf('/') + 1));
  };

  const handleChange: UploadProps['onChange'] = ({ fileList: newFileList }) =>
    setFileList(newFileList);

  // Custom upload request
  const customRequest: UploadProps['customRequest'] = async (options) => {
    const { file, onSuccess, onError } = options;
    try {
      // Call our API to upload the file
      const filename = await uploadImage(file as File);
      
      // We need to tell Antd Upload that it succeeded
      // and store the server's filename in the file object response
      onSuccess?.({ filename });
      
      // Also update fileList item status? 
      // Antd usually handles status updates via onChange if onSuccess is called.
    } catch (err) {
      console.error("Upload failed", err);
      onError?.(err as Error);
      message.error(`${(file as File).name} upload failed.`);
    }
  };

  const onFinish = async (values: any) => {
    setLoading(true);
    try {
      // Extract successful uploads filenames
      const imageFilenames = fileList
        .filter(file => file.status === 'done' && file.response?.filename)
        .map(file => file.response.filename);

      // Convert dayjs object to ISO string
      const payload: CreateMealRequest = {
        name: values.name,
        description: values.description,
        date: values.date ? values.date.toISOString() : new Date().toISOString(),
        image_filenames: imageFilenames,
      };
      
      await onSubmit(payload);
      
      // Reset form
      form.resetFields();
      setFileList([]);
      // Reset defaults
      form.setFieldsValue({ name: '好饭', date: dayjs() });

      message.success('Meal added successfully');
    } catch (error) {
      message.error('Failed to add meal');
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Form
        form={form}
        layout="vertical"
        onFinish={onFinish}
        autoComplete="off"
        initialValues={{ name: '好饭', date: dayjs() }}
      >
        <Form.Item
          label="Name"
          name="name"
          rules={[{ required: true, message: 'Please input the meal name!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Date"
          name="date"
          rules={[{ required: true, message: 'Please select a date!' }]}
        >
          <DatePicker style={{ width: '100%' }} />
        </Form.Item>

        <Form.Item
          label="Description"
          name="description"
        >
          <Input.TextArea />
        </Form.Item>

        <Form.Item label="Images">
          <Upload
            listType="picture-card"
            fileList={fileList}
            onPreview={handlePreview}
            onChange={handleChange}
            customRequest={customRequest}
          >
            {fileList.length >= 8 ? null : (
              <div>
                <PlusOutlined />
                <div style={{ marginTop: 8 }}>Upload</div>
              </div>
            )}
          </Upload>
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading}>
            Add Meal
          </Button>
        </Form.Item>
      </Form>

      <Modal open={previewOpen} title={previewTitle} footer={null} onCancel={handleCancel}>
        <img alt="example" style={{ width: '100%' }} src={previewImage} />
      </Modal>
    </>
  );
};

// Helper to get base64 for preview
const getBase64 = (file: File): Promise<string> =>
  new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = (error) => reject(error);
  });
