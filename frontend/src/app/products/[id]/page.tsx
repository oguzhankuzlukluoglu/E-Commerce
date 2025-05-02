import { useEffect, useState } from 'react'
import Image from 'next/image'
import { useAuth } from '@/store/authStore'
import { useCart } from '@/store/cartStore'
import { toast } from 'react-hot-toast'

interface Product {
  id: string
  name: string
  description: string
  price: number
  image: string
  category: string
  stock: number
}

export default function ProductDetailPage({ params }: { params: { id: string } }) {
  const { token } = useAuth()
  const { addItem } = useCart()
  const [product, setProduct] = useState<Product | null>(null)
  const [loading, setLoading] = useState(true)
  const [quantity, setQuantity] = useState(1)

  useEffect(() => {
    fetchProduct()
  }, [params.id, token])

  const fetchProduct = async () => {
    try {
      const response = await fetch(`/api/products/${params.id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })

      if (!response.ok) {
        throw new Error('Ürün yüklenemedi')
      }

      const data = await response.json()
      setProduct(data)
    } catch (error) {
      console.error('Error fetching product:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleAddToCart = () => {
    if (!product) return

    if (quantity > product.stock) {
      toast.error('Yeterli stok bulunmamaktadır')
      return
    }

    addItem({
      id: product.id,
      name: product.name,
      price: product.price,
      image: product.image,
      quantity,
    })

    toast.success('Ürün sepete eklendi')
  }

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

  if (!product) {
    return (
      <div className="bg-white">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col items-center justify-center py-24">
            <h1 className="text-2xl font-bold tracking-tight text-gray-900">Ürün bulunamadı</h1>
            <p className="mt-4 text-gray-500">
              Aradığınız ürün bulunamadı veya kaldırılmış olabilir.
            </p>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="bg-white">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-2xl py-16 sm:py-24 lg:max-w-none lg:py-32">
          <div className="lg:grid lg:grid-cols-2 lg:items-start lg:gap-x-8">
            <div className="flex flex-col">
              <div className="aspect-h-1 aspect-w-1 w-full">
                <Image
                  src={product.image}
                  alt={product.name}
                  className="h-full w-full object-cover object-center sm:rounded-lg"
                  width={500}
                  height={500}
                />
              </div>
            </div>

            <div className="mt-10 px-4 sm:mt-16 sm:px-0 lg:mt-0">
              <h1 className="text-3xl font-bold tracking-tight text-gray-900">{product.name}</h1>
              <div className="mt-3">
                <h2 className="sr-only">Ürün bilgileri</h2>
                <p className="text-3xl tracking-tight text-gray-900">{product.price} TL</p>
              </div>

              <div className="mt-6">
                <h3 className="sr-only">Açıklama</h3>
                <div className="space-y-6 text-base text-gray-700">{product.description}</div>
              </div>

              <div className="mt-6">
                <div className="flex items-center">
                  <p className="text-sm text-gray-500">Kategori:</p>
                  <p className="ml-2 text-sm font-medium text-gray-900">{product.category}</p>
                </div>
                <div className="mt-2 flex items-center">
                  <p className="text-sm text-gray-500">Stok:</p>
                  <p className="ml-2 text-sm font-medium text-gray-900">{product.stock} adet</p>
                </div>
              </div>

              <div className="mt-10 flex">
                <div className="mr-4">
                  <label htmlFor="quantity" className="sr-only">
                    Adet
                  </label>
                  <select
                    id="quantity"
                    name="quantity"
                    value={quantity}
                    onChange={(e) => setQuantity(Number(e.target.value))}
                    className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                  >
                    {[...Array(Math.min(10, product.stock))].map((_, i) => (
                      <option key={i + 1} value={i + 1}>
                        {i + 1}
                      </option>
                    ))}
                  </select>
                </div>
                <button
                  type="button"
                  onClick={handleAddToCart}
                  className="flex max-w-xs flex-1 items-center justify-center rounded-md border border-transparent bg-indigo-600 px-8 py-3 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:ring-offset-gray-50 sm:w-full"
                >
                  Sepete Ekle
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 