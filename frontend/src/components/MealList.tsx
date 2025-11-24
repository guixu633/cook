import React from 'react';
import { List, Card, Button, Popconfirm, Image, Typography } from 'antd';
import { DeleteOutlined } from '@ant-design/icons';
import type { Meal } from '../types/meal';

interface MealListProps {
  meals: Meal[];
  loading: boolean;
  onDelete: (id: number) => Promise<void>;
}

const { Text } = Typography;

export const MealList: React.FC<MealListProps> = ({ meals, loading, onDelete }) => {
  return (
    <List
      grid={{ gutter: 16, xs: 1, sm: 2, md: 3, lg: 3, xl: 4, xxl: 4 }}
      dataSource={meals}
      loading={loading}
      renderItem={(item) => (
        <List.Item>
          <Card
            title={item.name}
            size="small"
            extra={
              <Popconfirm
                title="Delete the meal"
                description="Are you sure to delete this meal?"
                onConfirm={() => onDelete(item.id)}
                okText="Yes"
                cancelText="No"
              >
                <Button type="text" danger icon={<DeleteOutlined />} />
              </Popconfirm>
            }
          >
            <p style={{ 
              marginBottom: '12px', 
              overflow: 'hidden', 
              textOverflow: 'ellipsis', 
              display: '-webkit-box', 
              WebkitLineClamp: 2, 
              WebkitBoxOrient: 'vertical' 
            }}>
              {item.description}
            </p>
            {item.image_urls && item.image_urls.length > 0 && (
              <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
                {item.image_urls.map((url, index) => (
                  <Image
                    key={index}
                    width={60}
                    height={60}
                    src={url}
                    style={{ objectFit: 'cover', borderRadius: '4px' }}
                    fallback="https://via.placeholder.com/60"
                  />
                ))}
              </div>
            )}
            <div style={{ marginTop: '10px' }}>
                <Text type="secondary" style={{ fontSize: '12px' }}>
                    {/* Show the user-entered date, fallback to created_at if missing (for old data) */}
                    {new Date(item.date || item.created_at).toLocaleDateString()}
                </Text>
            </div>
          </Card>
        </List.Item>
      )}
    />
  );
};
