// app/page.tsx
import Hero from "../../components/LandingPageComponents/Hero";
import Sidebar from "@/app/components/Sidebar";
import HowItWorks from "@/app/components/HowItWorks/HowItWorks";
import PopularSearches from "@/app/components/LandingPageComponents/PopularSearches";
import FeaturedDeals from "@/app/components/LandingPageComponents/FeaturedDeals";
import Footer from "@/app/components/LandingPageComponents/Footer";

export default function Home() {
  return (
    <main>
      {/* Sidebar fixed at top */}
      <Sidebar activePage="home" />

      {/* Sections with spacing */}
      <div className="space-y-30">
        <Hero />
        <HowItWorks />
        <PopularSearches />
        <FeaturedDeals />
      </div>

      <Footer />
    </main>
  );
}
