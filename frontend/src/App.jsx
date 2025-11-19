import React, { useEffect, useState } from 'react';
import MealList from './components/MealList';
import AddMeal from './components/AddMeal';
import { getMeals, deleteMeal } from './api';

function App() {
    const [meals, setMeals] = useState([]);

    const fetchMeals = async () => {
        try {
            const data = await getMeals();
            setMeals(data);
        } catch (error) {
            console.error('Error fetching meals:', error);
        }
    };

    useEffect(() => {
        fetchMeals();
    }, []);

    const handleMealAdded = (newMeal) => {
        setMeals([newMeal, ...meals]);
    };

    const handleDeleteMeal = async (id) => {
        if (window.confirm('Are you sure you want to delete this meal?')) {
            try {
                await deleteMeal(id);
                setMeals(meals.filter((meal) => meal.ID !== Number(id)));
            } catch (error) {
                console.error('Error deleting meal:', error);
                alert('Failed to delete meal: ' + error.message);
            }
        }
    };

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-4xl font-bold text-center mb-8">Daily Meal Tracker</h1>
            <AddMeal onMealAdded={handleMealAdded} />
            <MealList meals={meals} onDelete={handleDeleteMeal} />
        </div>
    );
}

export default App;
