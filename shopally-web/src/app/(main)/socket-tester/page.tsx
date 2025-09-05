"use client";

import { useEffect, useState } from "react";
import { io, Socket } from "socket.io-client";

interface PriceUpdate {
  id: string;
  newPriceETB: number;
  oldPriceETB: number;
}

export default function AlertsPage() {
  const [alerts, setAlerts] = useState<string[]>([]);
  const [updates, setUpdates] = useState<PriceUpdate[]>([]);
  const [socket, setSocket] = useState<Socket | null>(null);

  useEffect(() => {
    const socketInstance: Socket = io("http://localhost:4000");
    setSocket(socketInstance);

    socketInstance.on("connect", () => {
      console.log("🔌 Connected:", socketInstance.id);
    });

    socketInstance.on("alerts", (data: { productIds: string[] }) => {
      console.log("📢 Alerts received:", data.productIds);
      setAlerts(data.productIds);
    });

    socketInstance.on("priceUpdate", (update: PriceUpdate) => {
      console.log("💰 Price update:", update);
      setUpdates((prev) => [...prev, update]);
    });

    return () => {
      socketInstance.disconnect();
    };
  }, []);

  const createAlert = (productId: string) => {
    socket?.emit("createAlert", productId);
  };

  const deleteAlert = (productId: string) => {
    socket?.emit("deleteAlert", productId);
  };

  return (
    <div className="p-6 space-y-6">
      <h1 className="text-2xl font-bold">📢 Active Alerts</h1>
      <ul className="list-disc pl-6">
        {alerts.map((id) => (
          <li key={id} className="flex items-center gap-2">
            Product {id}
            <button
              onClick={() => deleteAlert(id)}
              className="text-sm text-red-500 hover:underline"
            >
              Delete
            </button>
          </li>
        ))}
      </ul>

      <div className="flex gap-2">
        <button
          onClick={() => createAlert("2")}
          className="px-4 py-2 bg-blue-600 text-white rounded"
        >
          Create Alert for Product 2
        </button>
      </div>

      <h2 className="text-xl font-semibold">💰 Price Updates</h2>
      <ul className="list-disc pl-6">
        {updates.map((u, idx) => (
          <li key={idx}>
            Product {u.id} dropped from {u.oldPriceETB} → {u.newPriceETB} ETB
          </li>
        ))}
      </ul>
    </div>
  );
}
