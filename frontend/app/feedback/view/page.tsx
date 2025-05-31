'use client';

import { useEffect, useState } from 'react';
import ParticleBg from '@/components/ParticleBg';

interface Feedback {
  _id: string;
  name: string;
  email: string;
  message: string;
  rating: number;
  createdAt: string;
}

export default function ViewFeedbackPage() {
  const [feedbacks, setFeedbacks] = useState<Feedback[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchFeedbacks = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/feedback/');
        if (!response.ok) {
          throw new Error('Failed to fetch feedbacks');
        }
        const data = await response.json();
        setFeedbacks(data);
      } catch (err) {
        setError('Failed to load feedbacks. Please try again later.');
        console.error('Error fetching feedbacks:', err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchFeedbacks();
  }, []);

  if (isLoading) {
    return (
      <div className="min-h-screen w-full flex items-center justify-center" style={{background: 'linear-gradient(135deg, #ff1e56 0%, #000000 100%)'}}>
        <div className="text-xl text-white drop-shadow-glow">Loading feedbacks...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen w-full flex items-center justify-center" style={{background: 'linear-gradient(135deg, #ff1e56 0%, #000000 100%)'}}>
        <div className="text-xl text-red-300 drop-shadow-glow">{error}</div>
      </div>
    );
  }

  return (
    <>
      <ParticleBg />
      <div className="min-h-screen w-full py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-4xl mx-auto">
          <h2 className="text-3xl font-extrabold text-center text-white mb-8 drop-shadow-glow">Feedback Entries</h2>
          <div className="space-y-8">
            {feedbacks.length === 0 ? (
              <p className="text-center text-red-200">No feedback entries yet.</p>
            ) : (
              feedbacks.map((feedback) => (
                <div
                  key={feedback._id}
                  className="bg-black/80 rounded-2xl shadow-2xl p-8 border border-red-700 backdrop-blur-md hover:shadow-pink-500/40 transition-shadow"
                  style={{boxShadow: '0 0 24px 2px #ff1e5688'}}
                >
                  <div className="flex justify-between items-start mb-4">
                    <div>
                      <h3 className="text-lg font-bold text-white drop-shadow-glow">{feedback.name}</h3>
                      <p className="text-sm text-red-200">{feedback.email}</p>
                    </div>
                    <div className="flex items-center">
                      <span className="text-yellow-400 mr-1 text-xl drop-shadow-glow">★</span>
                      <span className="text-white font-bold">{feedback.rating}/5</span>
                    </div>
                  </div>
                  <p className="text-white mb-4">{feedback.message}</p>
                  <p className="text-xs text-red-300">
                    Submitted on {new Date(feedback.createdAt).toLocaleDateString()}
                  </p>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </>
  );
} 