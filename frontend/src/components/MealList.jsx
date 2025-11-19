import React from 'react';

const MealList = ({ meals, onDelete }) => {
    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {meals.map((meal) => (
                <div key={meal.ID} className="border rounded-lg overflow-hidden shadow-lg">
                    {meal.image_url && (
                        <img src={meal.image_url} alt={meal.name} className="w-full h-48 object-cover" />
                    )}
                    <div className="p-4">
                        <h3 className="text-xl font-bold mb-2">{meal.name}</h3>
                        <p className="text-gray-600">{meal.date}</p>
                        <button
                            onClick={() => onDelete(meal.ID)}
                            className="mt-2 bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-3 rounded text-sm"
                        >
                            Delete
                        </button>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default MealList;
