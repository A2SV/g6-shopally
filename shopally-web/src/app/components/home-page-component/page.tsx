import React from "react";
import { Star } from "lucide-react";
import Image from "next/image";

interface CardComponentProps {
  mode: "dark" | "light";
}

// when you want to call component, call based on this
{
  /* <div className=" w-full max-w-3xl ">
  <div className="grid grid-cols-1 w-[85%]    lg:grid-cols-2 gap-4   h-fit p-4 max-w-full mt-4">
    <CardComponent mode={darkMode ? "dark" : "light"} />
    <CardComponent mode={darkMode ? "dark" : "light"} />
    <CardComponent mode={darkMode ? "dark" : "light"} />
    <CardComponent mode={darkMode ? "dark" : "light"} />
  </div>
</div> */
}

const CardComponent: React.FC<CardComponentProps> = ({ mode }) => {
  return (
    <div
      className={` w-full h-[308px] break-words border-[2px] rounded-[12px] ${
        mode === "dark"
          ? "border-[#262B32] bg-[#262B32]"
          : "border-[#E5E7EB] bg-[#FFFFFF]"
      }  `}
    >
      <div className="h-[40%] grid place-items-center rounded-tl-[11px] rounded-tr-[11px] bg-amber-300">
        <Image
          src="#"
          alt="Start Mixer"
          className=" justify-center text-2xl "
        />
      </div>
      <div className=" h-[60%]  p-4 ">
        <h6
          className={`text-[14px] ${
            mode === "light" ? "text-[#262B32]" : "text-[#FFFFFF]"
          } font-semibold`}
        >
          Professional Stand Mixer
        </h6>
        <div className=" justify-between flex items-center mt-3 mb-4">
          <h6 className=" text-[#FFD300] text-[18px] ">$89.69</h6>
          <div className=" flex items-center gap-1 ">
            <span className=" text-[12px] text-[#262B32]  ">4.8</span>
            <Star className=" w-3 h-3 text-yellow-400 " />
          </div>
        </div>
        <button className="bg-[#FFD300] w-full rounded-[5px] mb-2 h-[32px] ">
          Remove From Compare
        </button>
        <button className=" border-[2px] w-full rounded-[5px] h-[36px] text-[#FFD300] border-[#FFD300] ">
          Buy On Alibaba
        </button>
      </div>
    </div>
  );
};

export default CardComponent;
