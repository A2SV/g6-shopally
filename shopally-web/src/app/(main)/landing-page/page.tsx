// app/page.tsx
import Hero from "../../components/LandingPageComponents/Hero";
import Sidebar from "@/app/components/Sidebar";
import HowItWorks from "@/app/components/HowItWorks/HowItWorks";


export default function Home() {
  return (
    <main>
        <Sidebar activePage="home" />
      <Hero />
      <HowItWorks />
    </main>
  );
}
