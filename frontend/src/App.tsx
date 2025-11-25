import { useEffect, useState } from 'react';
import { Layout, Typography, Divider, message } from 'antd';
import { MealForm } from './components/MealForm';
import { MealList } from './components/MealList';
import { getMeals, createMeal, deleteMeal } from './api/meal';
import type { Meal, CreateMealRequest } from './types/meal';
import './App.css';

const { Header, Content, Footer } = Layout;
const { Title } = Typography;

function App() {
  const [meals, setMeals] = useState<Meal[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchMeals = async () => {
    setLoading(true);
    try {
      const data = await getMeals();
      setMeals(data);
    } catch (error) {
      message.error('Failed to load meals');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchMeals();
  }, []);

  const handleCreate = async (values: CreateMealRequest) => {
    await createMeal(values);
    await fetchMeals(); // Refresh list
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteMeal(id);
      message.success('Meal deleted');
      await fetchMeals(); // Refresh list
    } catch (error) {
      message.error('Failed to delete meal');
    }
  };

  return (
    <Layout className="layout" style={{ minHeight: '100vh' }}>
      <Header style={{ display: 'flex', alignItems: 'center', padding: '0 24px' }}>
        <div className="logo" style={{ color: 'white', fontSize: '20px', fontWeight: 'bold' }}>
          Cook App
        </div>
      </Header>
      <Content className="app-content">
        <div className="site-layout-content">
          <Title level={2}>Manage Meals</Title>
          
          <div style={{ marginBottom: '40px', maxWidth: '100%' }}>
            <Title level={4}>Add New Meal</Title>
            <MealForm onSubmit={handleCreate} />
          </div>

          <Divider />

          <Title level={4}>Meal List</Title>
          <MealList meals={meals} loading={loading} onDelete={handleDelete} />
        </div>
      </Content>
      <Footer style={{ textAlign: 'center' }}>Cook App ©{new Date().getFullYear()}</Footer>
    </Layout>
  );
}

export default App;
