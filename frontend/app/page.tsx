"use client";
import { motion } from "framer-motion";
import Link from "next/link";
import ParticleBg from '@/components/ParticleBg';

const MotionH1 = motion("h1");
const MotionP = motion("p");
const MotionDiv = motion("div");

export default function Home() {
  return (
    <>
      <ParticleBg />
      <section className="flex flex-col items-center justify-center min-h-[80vh] w-full relative z-10 text-center select-none">
        <MotionH1
          initial={{ opacity: 0, y: 40 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 1, ease: "easeOut" }}
          className="text-4xl sm:text-6xl md:text-7xl font-extrabold bg-gradient-to-r from-red-500 via-red-400 to-pink-500 bg-clip-text text-transparent drop-shadow-[0_0_24px_#ff2d2d] mb-6"
        >
          Share Your Thoughts.<br />
          <span className="text-red-400">Shape the Future.</span>
        </MotionH1>
        <MotionP
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.5, duration: 1 }}
          className="max-w-xl mx-auto text-lg sm:text-2xl text-gray-300 mb-10 font-medium drop-shadow-glow"
        >
          Your feedback drives innovation. Join us in building tomorrow&apos;s experience—one idea at a time.
        </MotionP>
        <MotionDiv
          initial={{ scale: 0.8, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          transition={{ delay: 1, type: "spring", stiffness: 200 }}
        >
          <Link
            href="/feedback"
            className="px-8 py-4 rounded-full bg-gradient-to-r from-red-600 to-pink-600 text-white text-xl font-bold shadow-lg hover:shadow-red-500/50 hover:scale-105 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-red-400 animate-pulse"
          >
            Give Feedback
          </Link>
        </MotionDiv>
        <MotionDiv
          initial={{ opacity: 0 }}
          animate={{ opacity: 0.5 }}
          transition={{ delay: 1.5, duration: 1 }}
          className="absolute bottom-8 left-1/2 -translate-x-1/2 w-full flex justify-center pointer-events-none"
        >
          <div className="h-32 w-32 bg-red-500/10 rounded-full blur-3xl animate-float" />
        </MotionDiv>
      </section>
    </>
  );
}
