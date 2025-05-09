'use client';

import { useCallback, useEffect, useState } from 'react';
import Particles, { initParticlesEngine } from '@tsparticles/react';
import type { Engine, Container } from '@tsparticles/engine';
import { loadSlim } from '@tsparticles/slim';

export default function ParticlesComponent() {
  const [init, setInit] = useState(false);

  useEffect(() => {
    const initEngine = async () => {
      await initParticlesEngine(async (engine: Engine) => {
        await loadSlim(engine);
      });
      setInit(true);
    };
    
    initEngine();
  }, []);

  const particlesLoaded = useCallback(async (container?: Container) => {
    // Optional: Do something with the container
    console.log('Particles container loaded', container);
  }, []);

  return (
    <>
      {init && (
        <Particles
          id="bg-canvas"
          className="bg-canvas"
          particlesLoaded={particlesLoaded}
          options={{
            background: { color: "transparent" },
            particles: {
              color: { value: "#ff2d2d" },
              links: { enable: true, color: "#ff2d2d", opacity: 0.5 },
              move: { enable: true, speed: 1 },
              number: { value: 50 },
              opacity: { value: 0.7 },
              size: { value: 2 }
            },
            fullScreen: { enable: false },
            detectRetina: true
          }}
        />
      )}
    </>
  );
} 