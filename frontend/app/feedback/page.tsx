'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import ParticleBg from '@/components/ParticleBg';

export default function FeedbackPage() {
  const router = useRouter();
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    message: '',
    rating: 0
  });
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/feedback/`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        router.push('/feedback/view');
      } else {
        throw new Error('Failed to submit feedback');
      }
    } catch (error) {
      console.error('Error submitting feedback:', error);
      alert('Failed to submit feedback. Please try again.');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <>
      <ParticleBg />
      <div className="min-h-screen w-full flex items-center justify-center">
        <div className="max-w-md w-full bg-black/80 rounded-2xl shadow-2xl p-10 border border-red-700 backdrop-blur-md" style={{boxShadow: '0 0 32px 4px #ff1e5688'}}>
          <h2 className="text-3xl font-extrabold text-center text-white mb-8 drop-shadow-glow">Submit Feedback</h2>
          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label htmlFor="name" className="block text-sm font-medium text-red-300">Name</label>
              <input
                type="text"
                id="name"
                required
                className="mt-1 block w-full rounded-md border border-red-700 bg-black/60 text-white shadow-inner focus:border-red-400 focus:ring-red-400"
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              />
            </div>
            <div>
              <label htmlFor="email" className="block text-sm font-medium text-red-300">Email</label>
              <input
                type="email"
                id="email"
                required
                className="mt-1 block w-full rounded-md border border-red-700 bg-black/60 text-white shadow-inner focus:border-red-400 focus:ring-red-400"
                value={formData.email}
                onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              />
            </div>
            <div>
              <label htmlFor="rating" className="block text-sm font-medium text-red-300">Rating</label>
              <select
                id="rating"
                required
                className="mt-1 block w-full rounded-md border border-red-700 bg-black/60 text-white shadow-inner focus:border-red-400 focus:ring-red-400"
                value={formData.rating}
                onChange={(e) => setFormData({ ...formData, rating: parseInt(e.target.value) })}
              >
                {[1, 2, 3, 4, 5].map((num) => (
                  <option key={num} value={num} className="bg-black text-white">
                    {num}
                  </option>
                ))}
              </select>
            </div>
            <div>
              <label htmlFor="message" className="block text-sm font-medium text-red-300">Message</label>
              <textarea
                id="message"
                required
                rows={4}
                className="mt-1 block w-full rounded-md border border-red-700 bg-black/60 text-white shadow-inner focus:border-red-400 focus:ring-red-400"
                value={formData.message}
                onChange={(e) => setFormData({ ...formData, message: e.target.value })}
              />
            </div>
            <button
              type="submit"
              disabled={isSubmitting}
              className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-lg text-lg font-bold text-white bg-gradient-to-r from-red-500 via-pink-500 to-red-700 hover:from-pink-600 hover:to-red-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 disabled:opacity-50 drop-shadow-glow"
              style={{boxShadow: '0 0 16px 2px #ff1e56cc'}}
            >
              {isSubmitting ? 'Submitting...' : 'Submit Feedback'}
            </button>
          </form>
        </div>
      </div>
    </>
  );
} 