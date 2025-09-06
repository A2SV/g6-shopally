import 'package:flutter/material.dart';
import 'package:shopallymobile/features/shopping_assistant/domain/entities/product_entity.dart';

class ProductScreen extends StatelessWidget {
  final List<ProductEntity> products;
  const ProductScreen({super.key, required this.products});

  @override
  Widget build(BuildContext context) {
    final cs = Theme.of(context).colorScheme;
    final textTheme = Theme.of(context).textTheme;

    return Scaffold(
      backgroundColor: cs.background, // Set background color for the whole screen
      appBar: AppBar(
        title: const Text("Products"),
        centerTitle: true,
        backgroundColor: cs.surface,
        elevation: 0,
        titleTextStyle:
            textTheme.headlineSmall?.copyWith(fontWeight: FontWeight.bold, color: cs.onSurface),
        leading: IconButton( // Added a leading icon for navigation if needed
          icon: Icon(Icons.arrow_back_ios_new_rounded, color: cs.onSurface),
          onPressed: () {
            // Navigator.pop(context); // Implement navigation back
          },
        ),
      ),
      body: LayoutBuilder(
        builder: (context, constraints) {
          final isWide = constraints.maxWidth > 900;
          final isMedium = constraints.maxWidth > 650;
          final crossAxisCount = isWide ? 3 : (isMedium ? 2 : 1);

          // Adjusted aspect ratio for better look and to prevent overflow in smaller cards
          // Taller cards for single column, wider for multiple to balance text and image
          final double childAspectRatio;
          if (crossAxisCount == 1) {
            childAspectRatio = 1.0; // Taller for single column
          } else if (crossAxisCount == 2) {
            childAspectRatio = 0.8; // A bit taller for two columns
          } else {
            childAspectRatio = 0.85; // Similar for three columns
          }


          return GridView.builder(
            padding: const EdgeInsets.all(16),
            itemCount: products.length,
            gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: crossAxisCount,
              mainAxisSpacing: 16,
              crossAxisSpacing: 16,
              childAspectRatio: childAspectRatio,
            ),
            itemBuilder: (context, index) {
              final p = products[index];
              return Card(
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(16),
                ),
                elevation: 4,
                shadowColor: cs.shadow.withOpacity(0.1),
                clipBehavior: Clip.antiAlias, // Ensures content respects rounded corners
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Expanded(
                      flex: 3,
                      child: Stack( // Use Stack for potential overlays like favorites
                        children: [
                          Positioned.fill(
                            child: Image.network(
                              p.imageUrl,
                              fit: BoxFit.cover,
                              errorBuilder: (_, __, ___) => Container(
                                color: cs.surfaceVariant,
                                alignment: Alignment.center,
                                child: Icon(Icons.broken_image_rounded, size: 48, color: cs.onSurfaceVariant.withOpacity(0.6)),
                              ),
                            ),
                          ),
                          // Example: Add a favorite button
                          Positioned(
                            top: 8,
                            right: 8,
                            child: Container(
                              decoration: BoxDecoration(
                                color: cs.surface.withOpacity(0.7),
                                shape: BoxShape.circle,
                              ),
                              child: IconButton(
                                icon: Icon(Icons.favorite_border_rounded, color: cs.primary),
                                onPressed: () {
                                  // Handle favorite tap
                                },
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    Expanded(
                      flex: 2,
                      child: Padding(
                        padding: const EdgeInsets.all(12),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              p.name,
                              maxLines: 2,
                              overflow: TextOverflow.ellipsis,
                              style: textTheme.titleMedium?.copyWith(
                                fontWeight: FontWeight.bold,
                                color: cs.onSurface,
                              ),
                            ),
                            const SizedBox(height: 8),
                            // Overflow Fix: Wrap the Row with Flexible or use a combination of Flexible/Expanded
                            // for internal elements. Here, we'll ensure price/stock has enough space
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween, // Distribute space
                              children: [
                                Expanded( // Allow rating/reviews to take available space
                                  child: Row(
                                    mainAxisSize: MainAxisSize.min, // Shrink to fit
                                    children: [
                                      Icon(Icons.star_rounded,
                                          size: 18, color: Colors.amber.shade600),
                                      const SizedBox(width: 4),
                                      Text(
                                        p.rating.toStringAsFixed(1),
                                        style: textTheme.bodyMedium?.copyWith(
                                            fontWeight: FontWeight.w600, color: cs.onSurface),
                                      ),
                                      const SizedBox(width: 4), // Added small space
                                      Expanded( // Let reviews text flexibly take space
                                        child: Text(
                                          " (${p.toString()})", // Added space before parenthesis
                                          maxLines: 1, // Ensure reviews don't wrap
                                          overflow: TextOverflow.ellipsis, // Truncate if too long
                                          style: textTheme.bodySmall?.copyWith(
                                              color: cs.onSurface.withOpacity(.6)),
                                        ),
                                      ),
                                    ],
                                  ),
                                ),
                                const SizedBox(width: 8), // Space between rating and badge
                                _StockBadge(inStock: p.inStock), // Stock badge should be compact
                              ],
                            ),
                            const SizedBox(height: 10),
                            Text(
                              "\$${p.price}",
                              style: textTheme.titleLarge?.copyWith(
                                fontWeight: FontWeight.w800,
                                color: cs.primary,
                              ),
                            ),
                          ],
                        ),
                      ),
                    ),
                  ],
                ),
              );
            },
          );
        },
      ),
    );
  }
}

class _StockBadge extends StatelessWidget {
  final bool inStock;
  const _StockBadge({required this.inStock});

  @override
  Widget build(BuildContext context) {
    final cs = Theme.of(context).colorScheme;
    final color = inStock ? Colors.green.shade600 : Colors.red.shade600;
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 4),
      decoration: BoxDecoration(
        color: color.withOpacity(.15),
        borderRadius: BorderRadius.circular(20),
        border: Border.all(color: color.withOpacity(0.3), width: 0.8),
      ),
      child: Text(
        inStock ? "In Stock" : "Out of Stock",
        style: TextStyle(fontSize: 12, color: color, fontWeight: FontWeight.w600),
      ),
    );
  }
}