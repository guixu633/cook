import axios from 'axios';
import type { CreateMealRequest, Meal } from '../types/meal';

// Use environment variable for base URL if available (for production build)
// In development (Vite), '/api' is proxied to the backend.
// In production build, if served separately, you might need to point to the full URL.
// But if frontend is built and served, usually we want relative path or configurable.
// For now, keeping '/api' works if Nginx reverse proxy is set up on server, 
// or if we want to hardcode the server IP for a standalone frontend build (not recommended for CORS reasons usually, but feasible).

// If you want the built frontend to talk to api.cook.guixuu.com directly without Nginx proxy on the same domain:
const baseURL = import.meta.env.PROD ? 'https://api.cook.guixuu.com:8080/api' : '/api';

const apiClient = axios.create({
  baseURL: baseURL,
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
