import 'package:flutter/material.dart';
import '../../domain/Entity/product_entity.dart';

class ProductDetail extends StatelessWidget {
  final ProductEntity product;
  const ProductDetail({required this.product, super.key});

  Widget _ratingStars(double rating) {
    return Row(
      children: List.generate(5, (index) {
        if (rating <= index) {
          return const Icon(Icons.star_border, color: Colors.grey, size: 20);
        } else if (index + 1 > rating && index < rating) {
          return const Icon(Icons.star_half, color: Colors.amber, size: 20);
        } else {
          return const Icon(Icons.star, color: Colors.amber, size: 20);
        }
      }),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Product Details'),
        backgroundColor: Colors.white,
        foregroundColor: Colors.black,
        elevation: 0,
        actions: [
          IconButton(icon: const Icon(Icons.share), onPressed: () {}),
          IconButton(icon: const Icon(Icons.favorite_border), onPressed: () {}),
        ],
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // 📸 Product Image
            ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: Image.network(
                product.imageUrl,
                height: 280,
                width: double.infinity,
                fit: BoxFit.cover,
              ),
            ),

            const SizedBox(height: 20),

            // 🏷️ Title
            Text(
              product.title,
              style: const TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              maxLines: 2,
              overflow: TextOverflow.ellipsis,
            ),

            const SizedBox(height: 10),

            // 💰 Price + Tax + Discount
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  "${product.price.usd} USD (${product.price.etb} ETB)",
                  style: const TextStyle(
                    fontSize: 20,
                    fontWeight: FontWeight.w700,
                    color: Colors.green,
                  ),
                ),
                if (product.discount > 0)
                  Text(
                    "-${product.discount}% OFF",
                    style: const TextStyle(
                      fontSize: 14,
                      fontWeight: FontWeight.bold,
                      color: Colors.red,
                    ),
                  ),
              ],
            ),
            if (product.tax > 0)
              Text(
                "Incl. tax: ${product.tax} USD",
                style: TextStyle(color: Colors.grey[600]),
              ),

            const SizedBox(height: 10),

            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Row(
                  children: [
                    _ratingStars(product.productRating),
                    const SizedBox(width: 8),
                    Text(
                      product.productRating.toStringAsFixed(1),
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        color: Colors.grey[700],
                      ),
                    ),
                  ],
                ),
              ],
            ),

            const SizedBox(height: 20),

            // 🚚 Delivery & Sold Count
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text("Delivery: ${product.deliveryEstimate}"),
                Text("Sold: ${product.numberSold}"),
              ],
            ),

            const SizedBox(height: 20),

            // 📌 Key Summary (bullets)
            if (product.summaryBullets.isNotEmpty) ...[
              const Text(
                "Key Features",
                style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),
              ),
              const SizedBox(height: 10),
              Column(
                children: product.summaryBullets.map((bullet) {
                  return Padding(
                    padding: const EdgeInsets.symmetric(vertical: 4),
                    child: Row(
                      children: [
                        const Icon(
                          Icons.check_circle,
                          size: 18,
                          color: Colors.green,
                        ),
                        const SizedBox(width: 8),
                        Expanded(
                          child: Text(
                            bullet,
                            style: TextStyle(color: Colors.grey[700]),
                          ),
                        ),
                      ],
                    ),
                  );
                }).toList(),
              ),
              const SizedBox(height: 20),
            ],

            // 📝 Description
            if (product.description.isNotEmpty) ...[
              const Text(
                "Description",
                style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),
              ),
              const SizedBox(height: 8),
              Text(
                product.description,
                style: TextStyle(height: 1.5, color: Colors.grey[700]),
              ),
              const SizedBox(height: 20),
            ],
            // 🤖 AI Match
            Container(
              padding: const EdgeInsets.all(16),
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  colors: [Colors.orange[100]!, Colors.orange[50]!],
                ),
                borderRadius: BorderRadius.circular(12),
              ),
              child: Row(
                children: [
                  Icon(Icons.auto_awesome, color: Colors.orange[800], size: 28),
                  const SizedBox(width: 12),
                  Expanded(
                    child: Text(
                      "${(product.aiMatchPercentage * 100).toStringAsFixed(0)}% match with your preferences",
                      style: TextStyle(
                        fontWeight: FontWeight.w600,
                        color: Colors.orange[900],
                      ),
                    ),
                  ),
                ],
              ),
            ),

            const SizedBox(height: 20),

            // 🔘 Actions
            Row(
              children: [
                Expanded(
                  child: OutlinedButton.icon(
                    onPressed: () => print("Open: ${product.deepLinkUrl}"),
                    icon: const Icon(Icons.open_in_new),
                    label: const Text("View Website"),
                  ),
                ),
                const SizedBox(width: 10),
                Expanded(
                  child: ElevatedButton.icon(
                    onPressed: () {},
                    icon: const Icon(Icons.shopping_cart),
                    label: const Text("Add to Cart"),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
