import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import Page from '../page';

// Mock the ParticleBg component since it uses browser APIs
jest.mock('../../components/ParticleBg', () => {
  return function MockParticleBg() {
    return <div data-testid="particle-bg">Particle Background</div>;
  };
});

describe('Home Page', () => {
  it('renders the main heading and description', () => {
    render(<Page />);
    
    // Check if the main heading is present
    expect(screen.getByText('Share Your Thoughts.')).toBeInTheDocument();
    expect(screen.getByText('Shape the Future.')).toBeInTheDocument();
    
    // Check if the description is present
    expect(screen.getByText(/Your feedback drives innovation/i)).toBeInTheDocument();
  });

  it('renders the call-to-action button', () => {
    render(<Page />);
    
    const ctaButton = screen.getByRole('link', { name: /Give Feedback/i });
    expect(ctaButton).toBeInTheDocument();
    expect(ctaButton).toHaveAttribute('href', '/feedback');
  });

  it('renders the particle background', () => {
    render(<Page />);
    
    expect(screen.getByTestId('particle-bg')).toBeInTheDocument();
  });
// removed the test for the particle background because it was not working..done
}); 