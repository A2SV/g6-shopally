import 'package:flutter/material.dart';

import '../../domain/entities/product_entity.dart';
import '../pages/products_list_page.dart';

class ProductLinkBox extends StatelessWidget {
  final String text;
  final List<ProductEntity> products;
  const ProductLinkBox({super.key, required this.products, required this.text});


  String cut(String text){
    if (text.length>30){
      return '${text.substring(0, 30)}...';
    }
    return text;
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Navigator.of(context).push(
          MaterialPageRoute(
            builder: (_) => ProductScreen(products: products),
          ),
        );
      },
      child: Container(
        padding: EdgeInsets.symmetric(horizontal: 16.0, vertical: 12.0),
        margin: EdgeInsets.symmetric(horizontal: 16.0),
        decoration: BoxDecoration(
          gradient: LinearGradient(
            colors: [
              Colors.white,
              Colors.grey.shade300,
            ],
              begin: Alignment.topLeft,
              end: Alignment.bottomRight,
              transform: GradientRotation(0.5),
          ),
          borderRadius: BorderRadius.all(Radius.circular(12.0)),
          boxShadow: [
            BoxShadow(
              color: Colors.black12,
              blurRadius: 8.0,
              offset: Offset(0, 2),
            ),
          ],
        ),
        child: Row (
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  children: [
                    const Icon(Icons.link, color: Colors.blue),
                    const SizedBox(width: 8.0),
                    Text(
                      'Product search',
      
                      style: TextStyle(
                        fontSize: 14.0,
                        fontWeight: FontWeight.w400,
                        color: const Color.fromARGB(221, 78, 77, 77),
                        fontFamily: 'Arial',
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 4.0),
                Text(
                  cut(text),
                  style: TextStyle(
                    fontSize: 14.0,
                    color: Colors.black87,
                    fontWeight: FontWeight.w500,
                    fontFamily: 'Arial',
                  ),
                ),
              ],
            ),
          Icon(Icons.arrow_forward_ios, color: Colors.black, size: 16.0),
          ],
        ),
      ),
    );
  }
}