import { useEffect, useState } from 'react'
import { useAuth } from '@/store/authStore'
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

export default function OrdersPage() {
  const { token } = useAuth()
  const [orders, setOrders] = useState<Order[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchOrders()
  }, [token])

  const fetchOrders = async () => {
    try {
      const response = await fetch('/api/orders', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })

      if (!response.ok) {
        throw new Error('Siparişler yüklenemedi')
      }

      const data = await response.json()
      setOrders(data)
    } catch (error) {
      console.error('Error fetching orders:', error)
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

  if (orders.length === 0) {
    return (
      <div className="bg-white">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col items-center justify-center py-24">
            <h1 className="text-2xl font-bold tracking-tight text-gray-900">Henüz Siparişiniz Yok</h1>
            <p className="mt-4 text-gray-500">
              Alışverişe başlamak için{' '}
              <Link href="/products" className="text-indigo-600 hover:text-indigo-500">
                ürünler sayfasını
              </Link>{' '}
              ziyaret edebilirsiniz.
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
          <h1 className="text-2xl font-bold tracking-tight text-gray-900">Siparişlerim</h1>

          <div className="mt-8">
            <div className="flow-root">
              <ul role="list" className="-my-6 divide-y divide-gray-200">
                {orders.map((order) => (
                  <li key={order.id} className="flex flex-col space-y-4 py-6">
                    <div className="flex items-center justify-between">
                      <div>
                        <p className="text-sm text-gray-500">
                          Sipariş #{order.id}
                        </p>
                        <p className="text-sm text-gray-500">
                          {new Date(order.createdAt).toLocaleDateString()}
                        </p>
                      </div>
                      <div>
                        <span
                          className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${
                            order.status === 'completed'
                              ? 'bg-green-100 text-green-800'
                              : order.status === 'processing'
                              ? 'bg-yellow-100 text-yellow-800'
                              : 'bg-gray-100 text-gray-800'
                          }`}
                        >
                          {order.status === 'completed'
                            ? 'Tamamlandı'
                            : order.status === 'processing'
                            ? 'İşleniyor'
                            : 'Beklemede'}
                        </span>
                      </div>
                    </div>

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

                    <div className="flex items-center justify-between border-t border-gray-200 pt-4">
                      <p className="text-base font-medium text-gray-900">Toplam</p>
                      <p className="text-base font-medium text-gray-900">{order.total} TL</p>
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 