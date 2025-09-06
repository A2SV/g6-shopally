import 'dart:core';

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/core/databasehelper/database_helper.dart';
import 'package:shopallymobile/features/saveditem/data/models/product_model.dart';
import 'package:shopallymobile/features/saveditem/domain/entities/product.dart';
import 'package:shopallymobile/features/saveditem/presentation/bloc/bloc/saved_product_bloc.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/savedpagewidget/price_alert.dart';
import 'package:shopallymobile/core/constants/const_color.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/savedpagewidget/rating.dart';
import 'package:shopallymobile/features/shopping_assistant/domain/entities/product_entity.dart';
import 'package:url_launcher/url_launcher.dart';
class ProductInfo extends StatefulWidget {
  final String id;
  final String title;
  final double price;
  final int minOrder;
  final double rating;
  final String image;
  final int issaved;
  final List<ProductEntity> compparedItems;
  final ProductEntity item;

  const ProductInfo({
    super.key,
    this.id = '1',
    this.title = 'Product Title',
    this.price = 99.99,
    this.minOrder = 2,
    this.rating = 4.8,
    this.image = 'asset/image/imge1.jpg',
    this.issaved = 0,
    required this.compparedItems,
    required this.item,
  });
  @override
  State<ProductInfo> createState() => _ProductInfoState();
}

class _ProductInfoState extends State<ProductInfo> {
  late Color _iconColor;
  bool isSaved = false;
  void openWebLink(String url) async {
    final Uri uri = Uri.parse(url);

    if (await canLaunchUrl(uri)) {
      await launchUrl(uri, mode: LaunchMode.externalApplication);
    } else {
      throw 'Could not launch $url';
    }
  }

  @override
  void initState() {
    super.initState();
    _checkIfSaved();
    print("isSaved: $isSaved");
  }

  Future<void> _checkIfSaved() async {
    final exists = await DatabaseHelper().idExists('Saveditems', widget.id);
    if (mounted) {
      setState(() {
        isSaved = exists;
      });
    }
  }
  bool isaddtocompare = false;
  Color compare = Colors.white;

  @override
  Widget build(BuildContext context) {
    _iconColor = isSaved ? Colors.red : Colors.white;
    

    return LayoutBuilder(
      builder: (context, constraints) {
        final screenWidth = MediaQuery.of(context).size.width;
        final isMobile = screenWidth < 600;
        final containerWidth = isMobile ? screenWidth * 0.45 : 300.0;
        final imageHeight = isMobile ? 140.0 : 200.0;
        final imageWidth = containerWidth;
        final titleFontSize = isMobile ? 15.0 : 16.0;
        final priceFontSize = isMobile ? 13.0 : 14.0;
        final minOrderFontSize = isMobile ? 11.0 : 12.0;

        return Center(
          child: Stack(
            children: [
              Container(
                width: containerWidth,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(15.0),
                  gradient: LinearGradient(
                    begin: Alignment.topLeft,
                    end: Alignment.bottomRight,
                    colors: [black.withOpacity(0.0), white.withOpacity(0.0)],
                  ),
                  boxShadow: [
                    BoxShadow(
                      color: const Color.fromARGB(
                        255,
                        137,
                        145,
                        148,
                      ).withOpacity(00),
                      blurRadius: 3.0,
                      spreadRadius: 2.0,
                      offset: const Offset(4, 4),
                    ),
                  ],
                ),
                child: Column(
                  mainAxisSize: MainAxisSize
                      .min, // Allow column to shrink-wrap its content
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Container(
                      height: imageHeight,
                      width: imageWidth,
                      decoration: BoxDecoration(
                        borderRadius: const BorderRadius.only(
                          topLeft: Radius.circular(15.0),
                          topRight: Radius.circular(15.0),
                        ),
                        image: DecorationImage(
                          image: NetworkImage(widget.image),
                          fit: BoxFit.cover,
                        ),
                      ),
                    ),
                    Expanded(
                      // child: SingleChildScrollView(
                      child: Padding(
                        padding: EdgeInsets.all(isMobile ? 0.0 : 10.0),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              widget.title.length > 15 
                              ? '${widget.title.substring(0, 15)}...' 
                              : widget.title,
                              maxLines: 1,
                              overflow: TextOverflow.ellipsis,
                              style: TextStyle(
                                fontSize: titleFontSize,
                                fontWeight: FontWeight.bold,
                                color: black,
                              ),
                            ),
                            Text(
                              '\$${widget.price}',
                              style: TextStyle(
                                fontSize: priceFontSize,
                                color: black,
                              ),
                            ),
                            Text(
                              'Min. order: ${widget.minOrder} pieces',
                              style: TextStyle(
                                fontSize: minOrderFontSize,
                                color: black,
                              ),
                            ),
                            Row(
                              children: [
                                // Icon(Icons.add_shopping_cart),
                                Text('Alibaba', style: TextStyle(color: black)),
                              ],
                            ),
                            Row(
                              children: [
                                Text(
                                  'Price Alert',
                                  style: TextStyle(
                                    fontSize: isMobile ? 12 : 14,
                                    color: black,
                                  ),
                                ),
                                PriceAlert(),
                              ],
                            ),
                            RatingBar(rating: widget.rating, size: 12),

                            Column(
                              mainAxisSize: MainAxisSize.min,
                              children: [
                                Row(
                                  children: [
                                    Expanded(
                                      child: ElevatedButton(
                                        style: ElevatedButton.styleFrom(
                                          backgroundColor: const Color.fromARGB(
                                            255,
                                            101,
                                            228,
                                            217,
                                          ).withOpacity(0.9),
                                          shape: RoundedRectangleBorder(
                                            borderRadius: BorderRadius.circular(
                                              8.0,
                                            ),
                                          ),
                                        ),
                                        onPressed: () async{
                                          final Uri url = Uri.parse(widget.item.deeplinkUrl);
                                          debugPrint('Launching URL: $url');
                                          if (await canLaunchUrl(url)) {
                                              await launchUrl(url, mode: LaunchMode.externalApplication);
                                            } else {
                                              debugPrint('Could not launch $url');
                                            }
                                        },
                                        child: Text(
                                          'Buy On AliExpress',
                                          textAlign: TextAlign.center,
                                          style: TextStyle(
                                            fontSize: isMobile ? 12 : 14,
                                            color: black,
                                          ),
                                        ),
                                      ),
                                    ),
                                  ],
                                ),
                              ],
                            ),
                          ],
                        ),
                      ),
                      // ),
                    ),
                  ],
                ),
              ),
              Positioned(
                top: 0,
                right: 0,
                child: IconButton(
                  icon: Icon(Icons.add_box, color: compare),
                  onPressed: () {
                    // TODO: Implement delete functionality
                    setState(() {
                      if (isaddtocompare == false) {
                        isaddtocompare = true;
                        compare = Colors.blue;
                        widget.compparedItems.add(widget.item);
                      } else {
                        widget.compparedItems.remove(widget.item);
                        isaddtocompare = false;
                        compare = Colors.white;
                      }
                    });
                  },
                ),
              ),
              Positioned(
                top: 90,
                right: 130,
                child: IconButton(
                        icon: Icon(Icons.favorite, color: _iconColor),
                        onPressed: () async {
                          try {
                            final Product product = Product(
                              id: widget.id,
                              title: widget.title,
                              price: widget.price,
                              minOrder: widget.minOrder,
                              rating: widget.rating,
                              imageUrl: widget.image,
                              issaved: isSaved ? 1 : 0,
                              iscompare: 0,
                            );
                            print('Product: $product');

                            setState(() {
                              if (isSaved) {
                                isSaved = false;
                                _iconColor =  Colors.white;

                                context.read<SavedProductBloc>().add(
                                  RemoveProductEvent(widget.id),
                                );
                              } else {
                                isSaved = true;
                                _iconColor = Colors.red;
                                context.read<SavedProductBloc>().add(
                                  SaveProductEvent(ProductModel.fromEntity(product)),
                                );
                              }
                            });
                          } catch (e) {
                            // if (kDebugMode) {
                            print(e.toString());
                            // }
                          }
                        },
                      ),
              ),
            ],
          ),
        );
      },
    );
  }
}
