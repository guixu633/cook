export interface Meal {
  id: number;
  created_at: string;
  updated_at: string;
  name: string;
  description: string;
  date: string; // ISO string
  image_urls: string[];
}

export interface CreateMealRequest {
  name: string;
  description: string;
  date: string; // ISO string
  image_urls: string[];
}
