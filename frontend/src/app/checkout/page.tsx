import { useEffect, useState } from 'react'
import { useCart } from '@/store/cartStore'
import { useAuth } from '@/store/authStore'
import { useRouter } from 'next/navigation'
import { toast } from 'react-hot-toast'
import Link from 'next/link'

interface Order {
  id: string
  items: {
    id: string
    name: string
    price: number
    quantity: number
  }[]
  total: number
  status: string
  createdAt: string
}

export default function CheckoutPage() {
  const { items, clearCart } = useCart()
  const { token } = useAuth()
  const router = useRouter()
  const [loading, setLoading] = useState(false)
  const [order, setOrder] = useState<Order | null>(null)

  const total = items.reduce((sum, item) => sum + item.price * item.quantity, 0)

  useEffect(() => {
    if (items.length === 0) {
      router.push('/cart')
    }
  }, [items, router])

  const handleCheckout = async () => {
    try {
      setLoading(true)

      const response = await fetch('/api/orders', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          items: items.map((item) => ({
            id: item.id,
            quantity: item.quantity,
          })),
        }),
      })

      if (!response.ok) {
        throw new Error('Sipariş oluşturulamadı')
      }

      const data = await response.json()
      setOrder(data)
      clearCart()
      toast.success('Siparişiniz başarıyla oluşturuldu')
    } catch (error) {
      console.error('Error creating order:', error)
      toast.error('Sipariş oluşturulurken bir hata oluştu')
    } finally {
      setLoading(false)
    }
  }

  if (order) {
    return (
      <div className="bg-white">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="mx-auto max-w-2xl py-16 sm:py-24 lg:max-w-none lg:py-32">
            <h1 className="text-2xl font-bold tracking-tight text-gray-900">Siparişiniz Alındı</h1>

            <div className="mt-8">
              <div className="flow-root">
                <ul role="list" className="-my-6 divide-y divide-gray-200">
                  {order.items.map((item) => (
                    <li key={item.id} className="flex py-6">
                      <div className="flex flex-1 flex-col">
                        <div>
                          <div className="flex justify-between text-base font-medium text-gray-900">
                            <h3>{item.name}</h3>
                            <p className="ml-4">{item.price} TL</p>
                          </div>
                          <p className="mt-1 text-sm text-gray-500">Adet: {item.quantity}</p>
                        </div>
                      </div>
                    </li>
                  ))}
                </ul>
              </div>

              <div className="mt-8 border-t border-gray-200 pt-8">
                <div className="flex items-center justify-between">
                  <p className="text-base font-medium text-gray-900">Toplam</p>
                  <p className="text-base font-medium text-gray-900">{order.total} TL</p>
                </div>
                <div className="mt-6">
                  <p className="text-sm text-gray-500">
                    Sipariş numaranız: {order.id}
                  </p>
                  <p className="mt-2 text-sm text-gray-500">
                    Sipariş durumu: {order.status}
                  </p>
                  <p className="mt-2 text-sm text-gray-500">
                    Sipariş tarihi: {new Date(order.createdAt).toLocaleDateString()}
                  </p>
                </div>
                <div className="mt-6">
                  <Link
                    href="/products"
                    className="w-full rounded-md border border-transparent bg-indigo-600 px-4 py-3 text-base font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:ring-offset-gray-50"
                  >
                    Alışverişe Devam Et
                  </Link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="bg-white">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-2xl py-16 sm:py-24 lg:max-w-none lg:py-32">
          <h1 className="text-2xl font-bold tracking-tight text-gray-900">Ödeme</h1>

          <div className="mt-8">
            <div className="flow-root">
              <ul role="list" className="-my-6 divide-y divide-gray-200">
                {items.map((item) => (
                  <li key={item.id} className="flex py-6">
                    <div className="flex flex-1 flex-col">
                      <div>
                        <div className="flex justify-between text-base font-medium text-gray-900">
                          <h3>{item.name}</h3>
                          <p className="ml-4">{item.price} TL</p>
                        </div>
                        <p className="mt-1 text-sm text-gray-500">Adet: {item.quantity}</p>
                      </div>
                    </div>
                  </li>
                ))}
              </ul>
            </div>

            <div className="mt-8 border-t border-gray-200 pt-8">
              <div className="flex items-center justify-between">
                <p className="text-base font-medium text-gray-900">Toplam</p>
                <p className="text-base font-medium text-gray-900">{total} TL</p>
              </div>
              <div className="mt-6">
                <button
                  onClick={handleCheckout}
                  disabled={loading}
                  className="w-full rounded-md border border-transparent bg-indigo-600 px-4 py-3 text-base font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:ring-offset-gray-50 disabled:opacity-50"
                >
                  {loading ? 'İşleniyor...' : 'Siparişi Tamamla'}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 