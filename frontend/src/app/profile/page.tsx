import { useEffect, useState } from 'react'
import { useAuth } from '@/store/authStore'
import { toast } from 'react-hot-toast'

interface User {
  id: number
  first_name: string
  last_name: string
  email: string
  role: string
  created_at: string
  updated_at: string
}

export default function ProfilePage() {
  const { token } = useAuth()
  const [user, setUser] = useState<User | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchUser()
  }, [token])

  const fetchUser = async () => {
    try {
      const response = await fetch('/api/auth/me', {
        headers: {
          Authorization: token ? `Bearer ${token}` : '',
        },
      })
      if (!response.ok) {
        throw new Error('Kullanıcı bilgileri yüklenemedi')
      }
      const data = await response.json()
      setUser(data)
    } catch (error) {
      setUser(null)
      toast.error('Kullanıcı bilgileri alınamadı')
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

  if (!user) {
    return (
      <div className="bg-white">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col items-center justify-center py-24">
            <h1 className="text-2xl font-bold tracking-tight text-gray-900">Kullanıcı Bulunamadı</h1>
            <p className="mt-4 text-gray-500">
              Kullanıcı bilgileriniz yüklenemedi. Lütfen tekrar giriş yapın.
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
          <h1 className="text-2xl font-bold tracking-tight text-gray-900">Profilim</h1>
          <div className="mt-8 space-y-6">
            <div>
              <span className="block text-sm font-medium text-gray-700">İsim</span>
              <div className="mt-1 text-lg text-gray-900">{user.first_name} {user.last_name}</div>
            </div>
            <div>
              <span className="block text-sm font-medium text-gray-700">E-posta</span>
              <div className="mt-1 text-lg text-gray-900">{user.email}</div>
            </div>
            <div>
              <span className="block text-sm font-medium text-gray-700">Rol</span>
              <div className="mt-1 text-lg text-gray-900">{user.role}</div>
            </div>
            <div>
              <span className="block text-sm font-medium text-gray-700">Kayıt Tarihi</span>
              <div className="mt-1 text-lg text-gray-900">{new Date(user.created_at).toLocaleDateString()}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 