import ParticleBg from '@/components/ParticleBg';

export default function AboutPage() {
  return (
    <>
      <ParticleBg />
      <div className="min-h-screen w-full flex flex-col items-center justify-center px-4 py-16">
        <h1 className="text-4xl font-extrabold text-white mb-6 drop-shadow-glow">About Us</h1>
        <p className="text-lg text-red-200 max-w-2xl text-center mb-10">
          We are passionate developers dedicated to building beautiful, modern web experiences. Our mission is to deliver high-quality, user-friendly applications that delight and inspire.
        </p>
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-8 max-w-4xl w-full">
          <div className="bg-black/80 rounded-xl border border-red-700 p-6 text-center shadow-lg">
            <h2 className="text-xl font-bold text-white mb-2">Arpit Srivastava</h2>
            <p className="text-red-200 mb-1">Frontend Developer</p>
            <p className="text-sm text-white">Loves crafting beautiful UIs and seamless user experiences.</p>
          </div>
          <div className="bg-black/80 rounded-xl border border-red-700 p-6 text-center shadow-lg">
            <h2 className="text-xl font-bold text-white mb-2">Arpit Srivastava</h2>
            <p className="text-red-200 mb-1">Backend Developer</p>
            <p className="text-sm text-white">Enjoys solving complex problems and optimizing performance.</p>
          </div>
          <div className="bg-black/80 rounded-xl border border-red-700 p-6 text-center shadow-lg">
            <h2 className="text-xl font-bold text-white mb-2">Arpit Srivastava</h2>
            <p className="text-red-200 mb-1">UI/UX Designer</p>
            <p className="text-sm text-white">Passionate about intuitive design and delightful interactions.</p>
          </div>
        </div>
      </div>
    </>
  );
} 