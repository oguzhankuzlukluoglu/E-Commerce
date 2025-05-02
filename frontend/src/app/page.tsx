import Link from 'next/link'
import Image from 'next/image'

const products = [
  {
    id: '1',
    name: 'Ürün 1',
    price: 99.99,
    image: 'https://images.unsplash.com/photo-1523275335684-37898b6baf30?ixlib=rb-1.2.1&auto=format&fit=crop&w=1567&q=80',
  },
  {
    id: '2',
    name: 'Ürün 2',
    price: 149.99,
    image: 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?ixlib=rb-1.2.1&auto=format&fit=crop&w=1567&q=80',
  },
  {
    id: '3',
    name: 'Ürün 3',
    price: 199.99,
    image: 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?ixlib=rb-1.2.1&auto=format&fit=crop&w=1567&q=80',
  },
]

const categories = [
  {
    name: 'Elektronik',
    href: '/products?category=electronics',
    image: 'https://images.unsplash.com/photo-1498049794561-7780e7231661?ixlib=rb-1.2.1&auto=format&fit=crop&w=1567&q=80',
  },
  {
    name: 'Giyim',
    href: '/products?category=clothing',
    image: 'https://images.unsplash.com/photo-1445205170230-053b83016050?ixlib=rb-1.2.1&auto=format&fit=crop&w=1567&q=80',
  },
  {
    name: 'Ev & Yaşam',
    href: '/products?category=home',
    image: 'https://images.unsplash.com/photo-1484101403633-562f891dc89a?ixlib=rb-1.2.1&auto=format&fit=crop&w=1567&q=80',
  },
]

export default function HomePage() {
  return (
    <main className="min-h-screen bg-gray-50 flex flex-col items-center justify-center px-4">
      <div className="max-w-2xl w-full text-center">
        <h1 className="text-4xl font-extrabold text-gray-900 mb-4">E-Commerce Sitesine Hoşgeldiniz!</h1>
        <p className="text-lg text-gray-600 mb-8">
          Modern ve hızlı bir alışveriş deneyimi için ürünlerimizi keşfedin. Go backend ile tam entegre, güvenli ve hızlı bir e-ticaret platformu!
        </p>
        <Link
          href="/products"
          className="inline-block rounded-md bg-indigo-600 px-8 py-3 text-lg font-semibold text-white shadow hover:bg-indigo-700 transition"
        >
          Ürünleri Görüntüle
        </Link>
      </div>
    </main>
  )
} 