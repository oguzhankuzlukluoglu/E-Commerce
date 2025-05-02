import { useEffect, useState } from 'react'
import Image from 'next/image'
import { useRouter } from 'next/navigation'

interface Product {
  id: number
  name: string
  description: string
  price: number
  image_url: string
  category: string
  stock: number
}

export default function ProductDetailPage({ params }: { params: { id: string } }) {
  const [product, setProduct] = useState<Product | null>(null)
  const [loading, setLoading] = useState(true)
  const router = useRouter()

  useEffect(() => {
    fetchProduct()
  }, [params.id])

  const fetchProduct = async () => {
    try {
      const response = await fetch(`/api/products/${params.id}`)
      if (!response.ok) {
        throw new Error('Ürün yüklenemedi')
      }
      const data = await response.json()
      setProduct(data)
    } catch (error) {
      setProduct(null)
    } finally {
      setLoading(false)
    }
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
                  src={product.image_url}
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
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 