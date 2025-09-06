import 'package:flutter/material.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';

class ProductDetailPage extends StatelessWidget {
  const ProductDetailPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title:  Text(getText('product_details')),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
        AspectRatio(
          aspectRatio: 1,
          child: Container(
            decoration: BoxDecoration(
          color: Colors.grey.shade200,
          borderRadius: BorderRadius.circular(12),
            ),
            child: const Icon(Icons.image, size: 80, color: Colors.grey),
          ),
        ),
        const SizedBox(height: 16),
         Text(
           getText("product_name"),
          style: TextStyle(fontSize: 22, fontWeight: FontWeight.bold),
        ),
        const SizedBox(height: 8),
        const Text(
          '\$99.99',
          style: TextStyle(fontSize: 18, color: Colors.green, fontWeight: FontWeight.w600),
        ),
        const SizedBox(height: 16),
         Text(
          getText('description'),
          style: TextStyle(fontSize: 16, fontWeight: FontWeight.w600),
        ),
        const SizedBox(height: 8),
        Text(
          getText('This is a placeholder description for the product. Replace with real data.'),
        ),
        const SizedBox(height: 24),
        SizedBox(
          width: double.infinity,
          child: ElevatedButton(
            onPressed: () {
          // TODO: Add to cart logic
            },
            child:  Text(getText('Add to Cart')),
          ),
        ),
          ],
        ),
      )
    );
  }
}