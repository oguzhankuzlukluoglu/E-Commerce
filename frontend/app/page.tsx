import Link from 'next/link'

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
