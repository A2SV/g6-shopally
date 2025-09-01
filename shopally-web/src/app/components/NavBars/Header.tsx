"use client";

import { signOut, useSession } from "next-auth/react";

const Header = () => {
  const { data: session } = useSession();

  return (
    <header className="flex justify-between items-center p-4 bg-gray-100 dark:bg-gray-900">
      <h1 className="text-xl font-bold">ShopAlly</h1>

      {session ? (
        <div className="flex items-center gap-4">
          {session.user?.image && (
            <img
              src={session.user.image}
              alt={session.user.name || "User"}
              className="w-8 h-8 rounded-full"
            />
          )}
          <span>{session.user?.name}</span>
          <button
            onClick={() => signOut({ callbackUrl: "/api/auth/signin" })}
            className="px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600"
          >
            Logout
          </button>
        </div>
      ) : null}
    </header>
  );
};

export default Header;
