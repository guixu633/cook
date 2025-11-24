import React, { useState } from 'react';
import { Form, Input, Button, Space, DatePicker, message } from 'antd';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import type { CreateMealRequest } from '../types/meal';
import dayjs from 'dayjs';

interface MealFormProps {
  onSubmit: (values: CreateMealRequest) => Promise<void>;
}

export const MealForm: React.FC<MealFormProps> = ({ onSubmit }) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const onFinish = async (values: any) => {
    setLoading(true);
    try {
      // Convert dayjs object to ISO string
      const payload: CreateMealRequest = {
        ...values,
        date: values.date ? values.date.toISOString() : new Date().toISOString(),
      };
      
      await onSubmit(payload);
      // Reset form but keep default values if needed
      form.resetFields(); 
      // Reset explicitly to ensure default name and current date are set back if we want
      form.setFieldsValue({ name: '好饭', date: dayjs() });

      message.success('Meal added successfully');
    } catch (error) {
      message.error('Failed to add meal');
    } finally {
      setLoading(false);
    }
  };

  return (
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

      <Form.List name="image_urls">
        {(fields, { add, remove }) => (
          <>
            {fields.map(({ key, name, ...restField }) => (
              <Space key={key} style={{ display: 'flex', marginBottom: 8 }} align="baseline">
                <Form.Item
                  {...restField}
                  name={[name]}
                  rules={[{ required: true, message: 'Missing URL' }]}
                >
                  <Input placeholder="Image URL" />
                </Form.Item>
                <MinusCircleOutlined onClick={() => remove(name)} />
              </Space>
            ))}
            <Form.Item>
              <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                Add Image URL
              </Button>
            </Form.Item>
          </>
        )}
      </Form.List>

      <Form.Item>
        <Button type="primary" htmlType="submit" loading={loading}>
          Add Meal
        </Button>
      </Form.Item>
    </Form>
  );
};
