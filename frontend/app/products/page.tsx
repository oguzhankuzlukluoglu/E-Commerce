import { useEffect, useState } from 'react'
import Image from 'next/image'
import Link from 'next/link'

interface Product {
  id: number
  name: string
  description: string
  price: number
  image_url: string
  category: string
  stock: number
}

export default function ProductsPage() {
  const [products, setProducts] = useState<Product[]>([])
  const [loading, setLoading] = useState(true)
  const [selectedCategory, setSelectedCategory] = useState<string>('all')
  const [categories, setCategories] = useState<string[]>([])
  const [searchQuery, setSearchQuery] = useState('')

  useEffect(() => {
    fetchProducts()
  }, [])

  const fetchProducts = async () => {
    try {
      const response = await fetch('/api/products')
      if (!response.ok) {
        throw new Error('Ürünler yüklenemedi')
      }
      const data = await response.json()
      setProducts(data.products)
      // Kategorileri çıkar
      const uniqueCategories = Array.from(new Set(data.products.map((product: Product) => product.category)))
      setCategories(uniqueCategories)
    } catch (error) {
      console.error('Error fetching products:', error)
    } finally {
      setLoading(false)
    }
  }

  const filteredProducts = products.filter((product) => {
    const matchesCategory = selectedCategory === 'all' || product.category === selectedCategory
    const matchesSearch = product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                         product.description.toLowerCase().includes(searchQuery.toLowerCase())
    return matchesCategory && matchesSearch
  })

  if (loading) {
    return (
      <div className="bg-white">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col items-center justify-center py-24">
            <div className="h-8 w-8 animate-spin rounded-full border-4 border-indigo-600 border-t-transparent" />
            <p className="mt-4 text-gray-500">Yükleniyor...</p>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="bg-white">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-2xl py-16 sm:py-24 lg:max-w-none lg:py-32">
          <h1 className="text-2xl font-bold tracking-tight text-gray-900">Ürünler</h1>

          <div className="mt-8 flex flex-col space-y-4 sm:flex-row sm:space-x-4 sm:space-y-0">
            <div className="flex-1">
              <input
                type="text"
                placeholder="Ürün ara..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
            <div className="w-full sm:w-48">
              <select
                value={selectedCategory}
                onChange={(e) => setSelectedCategory(e.target.value)}
                className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              >
                <option value="all">Tüm Kategoriler</option>
                {categories.map((category) => (
                  <option key={category} value={category}>
                    {category}
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div className="mt-8 grid grid-cols-1 gap-y-12 sm:grid-cols-2 sm:gap-x-6 lg:grid-cols-3 xl:gap-x-8">
            {filteredProducts.map((product) => (
              <div key={product.id} className="group relative">
                <div className="relative h-80 w-full overflow-hidden rounded-lg bg-white group-hover:opacity-75 sm:h-64">
                  <Image
                    src={product.image_url}
                    alt={product.name}
                    className="h-full w-full object-cover object-center"
                    width={300}
                    height={300}
                  />
                </div>
                <div className="mt-4 flex items-center justify-between">
                  <div>
                    <h3 className="text-sm font-medium text-gray-900">
                      <Link href={`/products/${product.id}`}>
                        <span aria-hidden="true" className="absolute inset-0" />
                        {product.name}
                      </Link>
                    </h3>
                    <p className="mt-1 text-sm text-gray-500">{product.category}</p>
                  </div>
                  <p className="text-sm font-medium text-gray-900">{product.price} TL</p>
                </div>
                <div className="mt-2">
                  <p className="text-sm text-gray-500 line-clamp-2">{product.description}</p>
                </div>
                <div className="mt-4">
                  <Link
                    href={`/products/${product.id}`}
                    className="text-sm font-medium text-indigo-600 hover:text-indigo-500"
                  >
                    Detayları Görüntüle
                  </Link>
                </div>
              </div>
            ))}
          </div>

          {filteredProducts.length === 0 && (
            <div className="mt-8 text-center">
              <p className="text-gray-500">Aradığınız kriterlere uygun ürün bulunamadı.</p>
            </div>
          )}
        </div>
      </div>
    </div>
  )
} 