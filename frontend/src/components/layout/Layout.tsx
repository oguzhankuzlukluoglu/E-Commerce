import { ReactNode } from 'react'
import Header from './Header'
import Footer from './Footer'
import { Toaster } from 'react-hot-toast'

interface LayoutProps {
  children: ReactNode
}

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="flex min-h-screen flex-col">
      <Header />
      <main className="flex-grow">{children}</main>
      <Footer />
      <Toaster position="top-right" />
    </div>
  )
} 