import React, { useState } from 'react';
import { addMeal } from '../api';

const AddMeal = ({ onMealAdded }) => {
    const [name, setName] = useState('');
    const [date, setDate] = useState('');
    const [image, setImage] = useState(null);
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        const formData = new FormData();
        formData.append('name', name);
        formData.append('date', date);
        if (image) {
            formData.append('image', image);
        }

        try {
            const newMeal = await addMeal(formData);
            onMealAdded(newMeal);
            setName('');
            setDate('');
            setImage(null);
        } catch (error) {
            console.error('Error adding meal:', error);
            alert('Failed to add meal');
        } finally {
            setLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="mb-8 p-4 border rounded-lg shadow-md bg-white">
            <h2 className="text-2xl font-bold mb-4">Add New Meal</h2>
            <div className="mb-4">
                <label className="block text-gray-700 text-sm font-bold mb-2">Meal Name</label>
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    required
                />
            </div>
            <div className="mb-4">
                <label className="block text-gray-700 text-sm font-bold mb-2">Date</label>
                <input
                    type="date"
                    value={date}
                    onChange={(e) => setDate(e.target.value)}
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    required
                />
            </div>
            <div className="mb-4">
                <label className="block text-gray-700 text-sm font-bold mb-2">Image</label>
                <input
                    type="file"
                    onChange={(e) => setImage(e.target.files[0])}
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    accept="image/*"
                />
            </div>
            <button
                type="submit"
                disabled={loading}
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline disabled:opacity-50"
            >
                {loading ? 'Adding...' : 'Add Meal'}
            </button>
        </form>
    );
};

export default AddMeal;
