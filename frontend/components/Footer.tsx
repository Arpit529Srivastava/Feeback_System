import { FaGithub, FaTwitter, FaLinkedin } from "react-icons/fa";

const socials = [
  { href: "https://github.com/", icon: <FaGithub />, label: "GitHub" },
  { href: "https://twitter.com/", icon: <FaTwitter />, label: "Twitter" },
  { href: "https://linkedin.com/", icon: <FaLinkedin />, label: "LinkedIn" },
];

export default function Footer() {
  return (
    <footer className="w-full bg-black/90 border-t border-red-900 py-4 flex flex-col items-center gap-2 mt-12">
      <div className="flex gap-6">
        {socials.map((s) => (
          <a
            key={s.label}
            href={s.href}
            target="_blank"
            rel="noopener noreferrer"
            aria-label={s.label}
            className="text-gray-400 text-2xl hover:text-red-500 transition drop-shadow-glow"
          >
            {s.icon}
          </a>
        ))}
      </div>
      <span className="text-xs text-gray-500">&copy; {new Date().getFullYear()} Futuristic Feedback System</span>
    </footer>
  );
} 