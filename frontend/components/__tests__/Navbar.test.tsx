import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import Navbar from '../Navbar';

describe('Navbar', () => {
  it('renders the navbar with all links', () => {
    render(<Navbar />);
    
    // Check if the logo is present
    expect(screen.getByText('Futurist')).toBeInTheDocument();
    
    // Check if all navigation links are present
    expect(screen.getByText('Home')).toBeInTheDocument();
    expect(screen.getByText('About')).toBeInTheDocument();
    expect(screen.getByText('Feedback')).toBeInTheDocument();
    expect(screen.getByText('View Feedbacks')).toBeInTheDocument();
  });

  it('has correct href attributes for navigation links', () => {
    render(<Navbar />);
    
    // Check if links have correct href attributes
    expect(screen.getByText('Home').closest('a')).toHaveAttribute('href', '/');
    expect(screen.getByText('About').closest('a')).toHaveAttribute('href', '/about');
    expect(screen.getByText('Feedback').closest('a')).toHaveAttribute('href', '/feedback');
    expect(screen.getByText('View Feedbacks').closest('a')).toHaveAttribute('href', '/feedback/view');
  });
}); 