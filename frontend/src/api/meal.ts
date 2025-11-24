import axios from 'axios';
import type { CreateMealRequest, Meal } from '../types/meal';

const apiClient = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getMeals = async (): Promise<Meal[]> => {
  const response = await apiClient.get<Meal[]>('/meals');
  return response.data;
};

export const createMeal = async (data: CreateMealRequest): Promise<Meal> => {
  const response = await apiClient.post<Meal>('/meals', data);
  return response.data;
};

export const deleteMeal = async (id: number): Promise<void> => {
  await apiClient.delete(`/meals/${id}`);
};

export const uploadImage = async (file: File): Promise<string> => {
  const formData = new FormData();
  formData.append('file', file);
  const response = await apiClient.post<{ filename: string }>('/upload', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
  return response.data.filename;
};
