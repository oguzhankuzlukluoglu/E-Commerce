import { useEffect, useState } from 'react'
import { useAuth } from '@/store/authStore'
import { useRouter } from 'next/navigation'
import Image from 'next/image'
import { toast } from 'react-hot-toast'

interface Order {
  id: string
  user: {
    id: string
    name: string
    email: string
  }
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

export default function AdminOrdersPage() {
  const { token } = useAuth()
  const router = useRouter()
  const [orders, setOrders] = useState<Order[]>([])
  const [loading, setLoading] = useState(true)
  const [showOrderModal, setShowOrderModal] = useState(false)
  const [selectedOrder, setSelectedOrder] = useState<Order | null>(null)

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

  const handleStatusChange = async (orderId: string, newStatus: string) => {
    try {
      const response = await fetch(`/api/orders/${orderId}`, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          status: newStatus,
        }),
      })

      if (!response.ok) {
        throw new Error('Sipariş durumu güncellenemedi')
      }

      const data = await response.json()
      setOrders((prev) =>
        prev.map((order) => (order.id === data.id ? data : order))
      )
      toast.success('Sipariş durumu başarıyla güncellendi')
    } catch (error) {
      console.error('Error updating order status:', error)
      toast.error('Sipariş durumu güncellenirken bir hata oluştu')
    }
  }

  const openOrderModal = (order: Order) => {
    setSelectedOrder(order)
    setShowOrderModal(true)
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

  return (
    <div className="bg-white">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-2xl py-16 sm:py-24 lg:max-w-none lg:py-32">
          <h1 className="text-2xl font-bold tracking-tight text-gray-900">Sipariş Yönetimi</h1>

          <div className="mt-8">
            <div className="flow-root">
              <div className="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                  <table className="min-w-full divide-y divide-gray-300">
                    <thead>
                      <tr>
                        <th
                          scope="col"
                          className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0"
                        >
                          Sipariş
                        </th>
                        <th
                          scope="col"
                          className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                        >
                          Müşteri
                        </th>
                        <th
                          scope="col"
                          className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                        >
                          Tarih
                        </th>
                        <th
                          scope="col"
                          className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                        >
                          Durum
                        </th>
                        <th
                          scope="col"
                          className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                        >
                          Toplam
                        </th>
                      </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-200">
                      {orders.map((order) => (
                        <tr key={order.id}>
                          <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm sm:pl-0">
                            <div className="font-medium text-gray-900">#{order.id}</div>
                            <div className="text-gray-500">
                              {order.items.length} ürün
                            </div>
                          </td>
                          <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                            <div className="text-gray-900">{order.user.name}</div>
                            <div className="text-gray-500">{order.user.email}</div>
                          </td>
                          <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                            {new Date(order.createdAt).toLocaleDateString()}
                          </td>
                          <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                            <select
                              value={order.status}
                              onChange={(e) => handleStatusChange(order.id, e.target.value)}
                              className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                            >
                              <option value="pending">Beklemede</option>
                              <option value="processing">İşleniyor</option>
                              <option value="completed">Tamamlandı</option>
                            </select>
                          </td>
                          <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                            {order.total} TL
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Order Details Modal */}
      {showOrderModal && selectedOrder && (
        <div className="fixed inset-0 z-10 overflow-y-auto">
          <div className="flex min-h-screen items-end justify-center px-4 pt-4 pb-20 text-center sm:block sm:p-0">
            <div className="fixed inset-0 transition-opacity" aria-hidden="true">
              <div className="absolute inset-0 bg-gray-500 opacity-75"></div>
            </div>
            <span className="hidden sm:inline-block sm:h-screen sm:align-middle" aria-hidden="true">
              &#8203;
            </span>
            <div className="inline-block transform overflow-hidden rounded-lg bg-white px-4 pt-5 pb-4 text-left align-bottom shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6 sm:align-middle">
              <div className="absolute top-0 right-0 hidden pt-4 pr-4 sm:block">
                <button
                  type="button"
                  className="rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                  onClick={() => setShowOrderModal(false)}
                >
                  <span className="sr-only">Kapat</span>
                  <svg
                    className="h-6 w-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    aria-hidden="true"
                  >
                    <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
              <div className="sm:flex sm:items-start">
                <div className="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                  <h3 className="text-lg font-medium leading-6 text-gray-900">
                    Sipariş #{selectedOrder.id}
                  </h3>
                  <div className="mt-4">
                    <div className="space-y-4">
                      <div>
                        <h4 className="text-sm font-medium text-gray-900">Müşteri Bilgileri</h4>
                        <div className="mt-2 text-sm text-gray-500">
                          <div>{selectedOrder.user.name}</div>
                          <div>{selectedOrder.user.email}</div>
                        </div>
                      </div>
                      <div>
                        <h4 className="text-sm font-medium text-gray-900">Sipariş Detayları</h4>
                        <div className="mt-2 space-y-2">
                          {selectedOrder.items.map((item) => (
                            <div key={item.id} className="flex items-center space-x-4">
                              <div className="h-10 w-10 flex-shrink-0">
                                <Image
                                  src={item.image}
                                  alt={item.name}
                                  className="h-10 w-10 rounded-full"
                                  width={40}
                                  height={40}
                                />
                              </div>
                              <div className="flex-1">
                                <div className="text-sm font-medium text-gray-900">{item.name}</div>
                                <div className="text-sm text-gray-500">
                                  {item.quantity} adet × {item.price} TL
                                </div>
                              </div>
                              <div className="text-sm font-medium text-gray-900">
                                {item.price * item.quantity} TL
                              </div>
                            </div>
                          ))}
                        </div>
                      </div>
                      <div>
                        <h4 className="text-sm font-medium text-gray-900">Sipariş Durumu</h4>
                        <div className="mt-2">
                          <select
                            value={selectedOrder.status}
                            onChange={(e) =>
                              handleStatusChange(selectedOrder.id, e.target.value)
                            }
                            className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                          >
                            <option value="pending">Beklemede</option>
                            <option value="processing">İşleniyor</option>
                            <option value="completed">Tamamlandı</option>
                          </select>
                        </div>
                      </div>
                      <div>
                        <h4 className="text-sm font-medium text-gray-900">Toplam</h4>
                        <div className="mt-2 text-lg font-medium text-gray-900">
                          {selectedOrder.total} TL
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  )
} 