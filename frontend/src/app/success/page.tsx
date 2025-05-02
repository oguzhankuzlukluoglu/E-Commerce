import Link from 'next/link'
import { useCart } from '@/store/cartStore'
import { useEffect } from 'react'

export default function SuccessPage() {
  const { clearCart } = useCart()

  useEffect(() => {
    clearCart()
  }, [clearCart])

  return (
    <div className="bg-white">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col items-center justify-center py-24">
          <div className="rounded-full bg-green-100 p-3">
            <svg
              className="h-6 w-6 text-green-600"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              aria-hidden="true"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M5 13l4 4L19 7"
              />
            </svg>
          </div>
          <h1 className="mt-4 text-2xl font-bold tracking-tight text-gray-900 sm:text-3xl">
            Siparişiniz Başarıyla Alındı
          </h1>
          <p className="mt-4 text-gray-500">
            Siparişiniz için teşekkür ederiz. Siparişinizle ilgili detayları e-posta adresinize
            gönderdik.
          </p>
          <div className="mt-10 flex items-center gap-x-6">
            <Link
              href="/products"
              className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              Alışverişe Devam Et
            </Link>
            <Link href="/orders" className="text-sm font-semibold leading-6 text-gray-900">
              Siparişlerim <span aria-hidden="true">→</span>
            </Link>
          </div>
        </div>
      </div>
    </div>
  )
} 