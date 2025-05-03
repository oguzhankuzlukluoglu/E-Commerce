import React, { useState } from 'react';
import Layout from '../../components/Layout';
import { productAPI } from '../../services/api';
import ProductCard from '../../components/ProductCard';
import { Product } from '../../types';

interface ProductsProps {
  initialProducts: Product[];
  totalProducts: number;
}

const Products: React.FC<ProductsProps> = ({ initialProducts, totalProducts }) => {
  const [products, setProducts] = useState(initialProducts);
  const [page, setPage] = useState(1);
  const [category, setCategory] = useState('');
  const [loading, setLoading] = useState(false);

  const handleLoadMore = async () => {
    setLoading(true);
    try {
      const newProducts = await productAPI.getAll({
        page: page + 1,
        limit: 12,
        category,
      });
      setProducts([...products, ...newProducts]);
      setPage(page + 1);
    } catch (error) {
      console.error('Error loading more products:', error);
    }
    setLoading(false);
  };

  const handleCategoryChange = async (newCategory: string) => {
    setLoading(true);
    setCategory(newCategory);
    setPage(1);
    try {
      const newProducts = await productAPI.getAll({
        page: 1,
        limit: 12,
        category: newCategory,
      });
      setProducts(newProducts);
    } catch (error) {
      console.error('Error filtering products:', error);
    }
    setLoading(false);
  };

  return (
    <Layout>
      <div className="max-w-7xl mx-auto px-4 py-8">
        <h1 className="text-3xl font-bold text-gray-800 mb-8">All Products</h1>
        
        {/* Category Filter */}
        <div className="mb-8">
          <div className="flex space-x-4">
            <button
              onClick={() => handleCategoryChange('')}
              className={`px-4 py-2 rounded-md ${
                category === ''
                  ? 'bg-indigo-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
            >
              All
            </button>
            <button
              onClick={() => handleCategoryChange('electronics')}
              className={`px-4 py-2 rounded-md ${
                category === 'electronics'
                  ? 'bg-indigo-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
            >
              Electronics
            </button>
            <button
              onClick={() => handleCategoryChange('clothing')}
              className={`px-4 py-2 rounded-md ${
                category === 'clothing'
                  ? 'bg-indigo-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
            >
              Clothing
            </button>
            <button
              onClick={() => handleCategoryChange('books')}
              className={`px-4 py-2 rounded-md ${
                category === 'books'
                  ? 'bg-indigo-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
            >
              Books
            </button>
          </div>
        </div>

        {/* Products Grid */}
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {products.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>

        {/* Load More Button */}
        {products.length < totalProducts && (
          <div className="mt-8 text-center">
            <button
              onClick={handleLoadMore}
              disabled={loading}
              className="bg-indigo-600 text-white px-6 py-2 rounded-md hover:bg-indigo-700 disabled:opacity-50"
            >
              {loading ? 'Loading...' : 'Load More'}
            </button>
          </div>
        )}
      </div>
    </Layout>
  );
};

export async function getServerSideProps() {
  try {
    const products = await productAPI.getAll({ limit: 12 });
    const totalProducts = 100; // This should come from the API
    return {
      props: {
        initialProducts: products,
        totalProducts,
      },
    };
  } catch (error) {
    console.error('Error fetching products:', error);
    return {
      props: {
        initialProducts: [],
        totalProducts: 0,
      },
    };
  }
}

export default Products; 