//mock-alerts-server.ts
import { createServer } from "http";
import { Server } from "socket.io";
// create HTTP server
const httpServer = createServer();
// create Socket.IO server
const io = new Server(httpServer, {
    cors: {
        origin: "*", // allow all origins for testing
    },
});
// mock alerts storage
let alerts = ["1", "3"];
io.on("connection", (socket) => {
    console.log("✅ Client connected:", socket.id);
    // send active alerts immediately
    socket.emit("alerts", { productIds: alerts });
    // simulate price update every 8s
    const interval = setInterval(() => {
        const updatedProduct = {
            id: "1",
            newPriceETB: 1400,
            oldPriceETB: 1650,
        };
        console.log("📉 Sending price update:", updatedProduct);
        socket.emit("priceUpdate", updatedProduct);
    }, 8000);
    // handle alert creation
    socket.on("createAlert", (productId) => {
        if (!alerts.includes(productId)) {
            alerts.push(productId);
            io.emit("alerts", { productIds: alerts });
        }
    });
    // handle alert deletion
    socket.on("deleteAlert", (productId) => {
        alerts = alerts.filter((id) => id !== productId);
        io.emit("alerts", { productIds: alerts });
    });
    socket.on("disconnect", () => {
        console.log("❌ Client disconnected:", socket.id);
        clearInterval(interval);
    });
});
// start server
const PORT = 4000;
httpServer.listen(PORT, () => {
    console.log(`🚀 Mock alerts server running on http://localhost:${PORT}`);
});
