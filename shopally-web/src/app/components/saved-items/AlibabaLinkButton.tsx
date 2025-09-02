import React from "react";

type Props = {
  deeplinkUrl: string;
};

const OpenLinkButton: React.FC<Props> = ({ deeplinkUrl }) => {
  const handleClick = () => {
    window.open(deeplinkUrl, "_blank", "noopener,noreferrer");
  };

  return (
    <button
      onClick={handleClick}
      className="flex-1 px-4 py-2 rounded-lg bg-[#0D2A4B] text-white hover:bg-[#133864] transition"
    >
      Buy on AliExpress
    </button>
  );
};

export default OpenLinkButton;
