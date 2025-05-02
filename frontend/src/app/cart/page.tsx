import { useCart } from '@/store/cartStore'
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { toast } from 'react-hot-toast'

export default function CartPage() {
  const { items, removeItem, updateQuantity } = useCart()
  const router = useRouter()

  const total = items.reduce((sum, item) => sum + item.price * item.quantity, 0)

  const handleCheckout = () => {
    if (items.length === 0) {
      toast.error('Sepetiniz boş')
      return
    }

    router.push('/checkout')
  }

  if (items.length === 0) {
    return (
      <div className="bg-white">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col items-center justify-center py-24">
            <h1 className="text-2xl font-bold tracking-tight text-gray-900">Sepetiniz Boş</h1>
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
          <h1 className="text-2xl font-bold tracking-tight text-gray-900">Sepetim</h1>

          <div className="mt-8">
            <div className="flow-root">
              <ul role="list" className="-my-6 divide-y divide-gray-200">
                {items.map((item) => (
                  <li key={item.id} className="flex py-6">
                    <div className="h-24 w-24 flex-shrink-0 overflow-hidden rounded-md border border-gray-200">
                      <Image
                        src={item.image}
                        alt={item.name}
                        className="h-full w-full object-cover object-center"
                        width={96}
                        height={96}
                      />
                    </div>

                    <div className="ml-4 flex flex-1 flex-col">
                      <div>
                        <div className="flex justify-between text-base font-medium text-gray-900">
                          <h3>{item.name}</h3>
                          <p className="ml-4">{item.price} TL</p>
                        </div>
                      </div>
                      <div className="flex flex-1 items-end justify-between text-sm">
                        <div className="flex items-center">
                          <label htmlFor={`quantity-${item.id}`} className="mr-2">
                            Adet:
                          </label>
                          <select
                            id={`quantity-${item.id}`}
                            value={item.quantity}
                            onChange={(e) => updateQuantity(item.id, Number(e.target.value))}
                            className="rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                          >
                            {[...Array(10)].map((_, i) => (
                              <option key={i + 1} value={i + 1}>
                                {i + 1}
                              </option>
                            ))}
                          </select>
                        </div>

                        <button
                          type="button"
                          onClick={() => removeItem(item.id)}
                          className="font-medium text-indigo-600 hover:text-indigo-500"
                        >
                          Kaldır
                        </button>
                      </div>
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          </div>

          <div className="mt-8 border-t border-gray-200 pt-8">
            <div className="flex items-center justify-between">
              <p className="text-base font-medium text-gray-900">Toplam</p>
              <p className="text-base font-medium text-gray-900">{total} TL</p>
            </div>
            <div className="mt-6">
              <button
                onClick={handleCheckout}
                className="w-full rounded-md border border-transparent bg-indigo-600 px-4 py-3 text-base font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:ring-offset-gray-50"
              >
                Ödemeye Geç
              </button>
            </div>
            <div className="mt-6 flex justify-center text-center text-sm text-gray-500">
              <p>
                veya{' '}
                <Link href="/products" className="font-medium text-indigo-600 hover:text-indigo-500">
                  Alışverişe Devam Et
                  <span aria-hidden="true"> &rarr;</span>
                </Link>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 