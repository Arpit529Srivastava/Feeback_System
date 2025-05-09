import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import ParticleBg from "../components/ParticleBg";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Futuristic Feedback System",
  description: "Share your thoughts. Shape the future.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased bg-black min-h-screen relative`}>
        <ParticleBg />
        <Navbar />
        <main className="relative z-10 pt-16 min-h-[80vh]">{children}</main>
        <Footer />
      </body>
    </html>
  );
}
