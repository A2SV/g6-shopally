import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/features/saveditem/presentation/bloc/bloc/saved_product_bloc.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/savedpagewidget/price_alert.dart';
import 'package:shopallymobile/core/constants/const_color.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/savedpagewidget/rating.dart';

class SingleProductInfo extends StatefulWidget {
  const SingleProductInfo({super.key,
    required this.id,
    required this.title,
    required this.price,
    required this.minOrder,
    required this.rating,
    required this.image,
  });
  final String id;
  final String title;
  final double price;
  final int minOrder;
  final double rating;
  final String image;

  @override
  State<SingleProductInfo> createState() => _SingleProductInfoState();
}

class _SingleProductInfoState extends State<SingleProductInfo> {
  @override
  Widget build(BuildContext context) {
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
                      color: const Color.fromARGB(255, 137, 145, 148).withOpacity(00),
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
                        borderRadius: BorderRadius.only(
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
                              widget.title,
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
                                Text(
                                  'Alibaba',
                                  style: TextStyle(color: black),
                                ),
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
                            RatingBar(rating: widget.rating , size: 12,),

                            Column(
                              mainAxisSize: MainAxisSize.min,
                              children: [
                                Row(
                                  children: [
                                    Expanded(
                                      child: ElevatedButton(
                                        style: ElevatedButton.styleFrom(
                                          backgroundColor: const Color.fromARGB(255, 101, 228, 217).withOpacity(0.9),
                                          shape: RoundedRectangleBorder(
                                            borderRadius: BorderRadius.circular(8.0),
                                          ),
                                        ),
                                        onPressed: () {},
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
                  icon: const Icon(Icons.delete, color: Colors.white),
                  onPressed: () {
                    context.read<SavedProductBloc>().add(
                                  RemoveProductEvent(widget.id),
                                );
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
