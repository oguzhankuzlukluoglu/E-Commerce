import axios from 'axios';
import { User, Product, Order, Payment } from '../types';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add token to requests if it exists
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Auth API
export const authAPI = {
  login: async (email: string, password: string) => {
    const response = await api.post('/auth/login', { email, password });
    return response.data;
  },
  register: async (username: string, email: string, password: string) => {
    const response = await api.post('/auth/register', { username, email, password });
    return response.data;
  },
  getProfile: async () => {
    const response = await api.get('/auth/profile');
    return response.data;
  },
};

// Product API
export const productAPI = {
  getAll: async (params?: { page?: number; limit?: number; category?: string }) => {
    const response = await api.get('/products', { params });
    return response.data;
  },
  getById: async (id: number) => {
    const response = await api.get(`/products/${id}`);
    return response.data;
  },
  create: async (product: Partial<Product>) => {
    const response = await api.post('/products', product);
    return response.data;
  },
  update: async (id: number, product: Partial<Product>) => {
    const response = await api.put(`/products/${id}`, product);
    return response.data;
  },
  delete: async (id: number) => {
    const response = await api.delete(`/products/${id}`);
    return response.data;
  },
};

// Order API
export const orderAPI = {
  create: async (order: { items: { productId: number; quantity: number }[] }) => {
    const response = await api.post('/orders', order);
    return response.data;
  },
  getAll: async () => {
    const response = await api.get('/orders');
    return response.data;
  },
  getById: async (id: number) => {
    const response = await api.get(`/orders/${id}`);
    return response.data;
  },
  cancel: async (id: number) => {
    const response = await api.put(`/orders/${id}/cancel`);
    return response.data;
  },
};

// Payment API
export const paymentAPI = {
  create: async (payment: { orderId: number; amount: number; paymentMethod: string }) => {
    const response = await api.post('/payments', payment);
    return response.data;
  },
  getById: async (id: number) => {
    const response = await api.get(`/payments/${id}`);
    return response.data;
  },
  refund: async (id: number) => {
    const response = await api.post(`/payments/${id}/refund`);
    return response.data;
  },
};

// User API
export const userAPI = {
  updateProfile: async (user: Partial<User>) => {
    const response = await api.put('/users/profile', user);
    return response.data;
  },
  changePassword: async (currentPassword: string, newPassword: string) => {
    const response = await api.put('/users/password', { currentPassword, newPassword });
    return response.data;
  },
};

export default api; 