'use client';
import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Navbar() {
  const [user, setUser] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    // Token varsa kullanıcı bilgisini çek
    const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;
    if (token) {
      // Token'dan kullanıcı adını almak için örnek bir çözüm
      // Gerçek uygulamada backend'den /auth/user ile çekmek daha güvenli
      const payload = token.split('.')[1];
      try {
        const decoded = JSON.parse(atob(payload));
        setUser(decoded.name || decoded.email || 'Kullanıcı');
      } catch {
        setUser('Kullanıcı');
      }
    } else {
      setUser(null);
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    router.push('/auth/login');
  };

  return (
    <nav className="w-full bg-gray-900 text-white px-4 py-3 flex items-center justify-between">
      <span className="font-bold text-lg">E-Commerce</span>
      {user ? (
        <div className="flex items-center gap-4">
          <span>Hoşgeldin, {user}</span>
          <button onClick={handleLogout} className="bg-red-500 hover:bg-red-600 px-3 py-1 rounded text-sm">Çıkış Yap</button>
        </div>
      ) : (
        <div className="flex gap-2">
          <a href="/auth/login" className="hover:underline">Giriş Yap</a>
          <a href="/auth/register" className="hover:underline">Kayıt Ol</a>
        </div>
      )}
    </nav>
  );
} 