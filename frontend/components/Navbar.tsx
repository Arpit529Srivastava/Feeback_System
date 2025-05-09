"use client";
import Link from "next/link";
import { motion } from "framer-motion";
import Image from "next/image";

const navLinks = [
  { name: "Home", href: "/" },
  { name: "About", href: "/about" },
  { name: "Feedback", href: "/feedback" },
  { name: "View Feedbacks", href: "/feedback/view" },
];

export default function Navbar() {
  return (
    <nav className="fixed top-0 left-0 w-full z-50 bg-black/80 backdrop-blur border-b border-red-900 shadow-lg">
      <div className="max-w-6xl mx-auto flex items-center justify-between px-4 py-2">
        <Link href="/" className="flex items-center gap-2">
          <Image src="/file.svg" alt="Logo" width={32} height={32} className="drop-shadow-glow" />
          <span className="text-lg font-bold text-white tracking-widest hidden sm:inline">Futurist</span>
        </Link>
        <div className="flex gap-2 sm:gap-6">
          {navLinks.map((link) => (
            <motion.div
              key={link.name}
              whileHover={{ scale: 1.1, textShadow: "0 0 8px #ff2d2d" }}
              transition={{ type: "spring", stiffness: 300 }}
            >
              <Link
                href={link.href}
                className="text-gray-200 px-3 py-1 rounded transition-all duration-200 hover:text-red-400 focus:text-red-400 focus:outline-none"
              >
                {link.name}
              </Link>
            </motion.div>
          ))}
        </div>
      </div>
    </nav>
  );
} 