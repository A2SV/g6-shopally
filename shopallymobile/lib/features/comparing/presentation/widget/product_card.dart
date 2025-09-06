import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart';
import 'package:shopallymobile/features/comparing/presentation/widget/product_detail.dart';

import '../bloc/compare_bloc.dart';

class ProductCard extends StatelessWidget {
  final List<ProductEntity> productEntities;
  const ProductCard({super.key, required this.productEntities});

  Card _productWidget(ProductEntity product, BuildContext context) {
    return Card(
      child: InkWell(
        onTap: () {
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => ProductDetail(product: product),
            ),
          );
        },
        child: Padding(
          padding: const EdgeInsets.all(13.0),
          child: Column(
            spacing: 5,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                width: double.infinity,
                height: 120,
                decoration: BoxDecoration(
                  image: DecorationImage(
                    image: NetworkImage(product.imageUrl),
                    fit: BoxFit.cover,
                  ),
                ),
              ),
              Text(
                product.title,
                overflow: TextOverflow.ellipsis,
                maxLines: 1,
                style: TextStyle(
                  fontSize: 14,
                  fontWeight: FontWeight.w600,
                  color: Color.fromRGBO(38, 43, 50, 1),
                ),
              ),
              Text(
                '${product.price.etb} ETB',
                style: TextStyle(
                  // color: Color.fromRGBO(255, 211, 0, 1),
                  fontSize: 12,
                  fontWeight: FontWeight.w700,
                ),
              ),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Color.fromRGBO(255, 211, 0, 1),
                  ),
                  onPressed: () {
                    context.read<CompareBloc>().add(
                      RemoveProductFromComparisonEvent(productId: product.id),
                    );
                  },
                  child: Text('Remove'),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Column(
        children: [
          GridView.builder(
            shrinkWrap: true,
            physics: NeverScrollableScrollPhysics(),
            gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 2, // Two columns
              crossAxisSpacing: 3,
              mainAxisSpacing: 3,
              childAspectRatio: 0.7, // Adjust aspect ratio for better layout
            ),
            itemCount: productEntities.length,
            itemBuilder: (context, index) {
              return _productWidget(productEntities[index], context);
            },
          ),
          SizedBox(height: 75),
          SizedBox(
            width:
                double.infinity, // makes the button expand to full card width
            child: ElevatedButton(
              style: ElevatedButton.styleFrom(
                backgroundColor: Color.fromRGBO(255, 211, 0, 1),
                foregroundColor: Colors.black, // text color
                padding: EdgeInsets.symmetric(vertical: 14), // taller button
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(12),
                ),
                elevation: 3, // subtle shadow
              ),
              onPressed: () {
                context.read<CompareBloc>().add(
                  CompareProductsEvent(productEntities),
                );
              },
              child: Text(
                'Compare',
                style: TextStyle(fontWeight: FontWeight.bold, fontSize: 16),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
