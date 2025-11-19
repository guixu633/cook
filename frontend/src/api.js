const API_URL = import.meta.env.PROD ? '/api' : 'http://127.0.0.1:3000/api';

export const getMeals = async () => {
    const response = await fetch(`${API_URL}/meals`);
    const data = await response.json();
    return data.data;
};

export const addMeal = async (formData) => {
    const response = await fetch(`${API_URL}/meals`, {
        method: 'POST',
        body: formData,
    });
    const data = await response.json();
    return data.data;
};

export const deleteMeal = async (id) => {
    const response = await fetch(`${API_URL}/meals/${id}`, {
        method: 'DELETE',
    });
    if (!response.ok) {
        throw new Error('Failed to delete meal');
    }
    return await response.json();
};
