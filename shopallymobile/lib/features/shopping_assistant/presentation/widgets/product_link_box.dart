import 'package:flutter/material.dart';
import 'package:shopallymobile/features/saveditem/presentation/pages/detail_page.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/productlist/buttom_nav.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/productlist/product_info.dart';

import '../../domain/entities/product_entity.dart';
import '../pages/products_list_page.dart';

class ProductLinkBox extends StatefulWidget {
  final String text;
  final List<ProductEntity> products;
  const ProductLinkBox({super.key, required this.products, required this.text});

  @override
  State<ProductLinkBox> createState() => _ProductLinkBoxState();
}

class _ProductLinkBoxState extends State<ProductLinkBox> {
  String cut(String text) {
    if (text.length > 30) {
      return '${text.substring(0, 30)}...';
    }
    return text;
  }
  final List<ProductEntity> _savedItems = [];

  void initState() {
    super.initState();
    _savedItems.addAll(widget.products);
  }
   final List<ProductEntity> compparedItems = [];

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        // Navigator.of(context).push(
        //   MaterialPageRoute(
        //     builder: (_) => ,
        //   ),
        // );
     /////////////////////////////////////////////
              showModalBottomSheet(
              context: context,
              isScrollControlled: true,
              builder: (BuildContext context) {
                return Padding(
                  padding: const EdgeInsets.only(top: 15.0),
                  child: DraggableScrollableSheet(
                    expand: false,
                    builder: (context, scrollController) {
                      return Stack(
                        alignment: Alignment.topCenter,
                        children: [
                          Container(
                            margin: const EdgeInsets.only(top: 5),
                            decoration: const BoxDecoration(
                              color: Colors.white,
                              borderRadius: BorderRadius.only(
                                topLeft: Radius.circular(20),
                                topRight: Radius.circular(20),
                              ),
                            ),
                            child: CustomScrollView(
                              controller: scrollController,
                              slivers: <Widget>[
                                SliverAppBar(
                                  pinned: true,
                                  backgroundColor: Colors.white,
                                  automaticallyImplyLeading: false,
                                  elevation: 0,
                                  title: Row(
                                    mainAxisAlignment:
                                        MainAxisAlignment.spaceBetween,
                                    children: [
                                      IconButton(
                                        icon: const Icon(
                                          Icons.arrow_back_ios,
                                          color: Colors.black,
                                        ),
                                        onPressed: () {
                                          Navigator.of(context).pop();
                                        },
                                      ),
                                      const Text(
                                        'Product Search',
                                        style: TextStyle(
                                          fontSize: 18,
                                          fontWeight: FontWeight.bold,
                                          color: Colors.black,
                                        ),
                                      ),
                                      TextButton(
                                        onPressed: () {
                                          debugPrint('this all is compared products ${compparedItems}');
                                        },
                                        child: const Text('Compare'),
                                      ),
                                    ],
                                  ),
                                  bottom: PreferredSize(
                                    preferredSize: const Size.fromHeight(40),
                                    child: Align(
                                      alignment: Alignment.centerLeft,
                                      child: Padding(
                                        padding: const EdgeInsets.only(left: 16.0, bottom: 8.0),
                                        child: Text(
                                          '',
                                          style: Theme.of(context)
                                              .textTheme
                                              .titleLarge
                                              ?.copyWith(
                                                fontWeight: FontWeight.bold,
                                                color: Colors.black,
                                              ),
                                        ),
                                      ),
                                    ),
                                  ),
                                ),
                                SliverToBoxAdapter(
                                  child: Padding(
                                    padding: const EdgeInsets.all(16.0),
                                    child: Column(
                                      crossAxisAlignment:
                                          CrossAxisAlignment.start,
                                      children:  [
                                        SizedBox(height: 8),
                                        Text(
                                          'Search results for ${widget.text}',
                                        ),
                                      ],
                                    ),
                                  ),
                                ),
                                 SliverPadding(
                  padding: const EdgeInsets.all(8.0),
                  sliver: SliverGrid(
                    gridDelegate:
                        const SliverGridDelegateWithFixedCrossAxisCount(
                          crossAxisCount: 2,
                          crossAxisSpacing: 2.0,
                          mainAxisSpacing: 8.0,
                          childAspectRatio: 0.55,
                        ),
                    delegate: SliverChildBuilderDelegate((context, index) {
                      final item = _savedItems[index];
                      return GestureDetector(
                        onTap: () {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => DetailPage(item: item),
                            ),
                          );
                        },

                        child: ProductInfo(id: item.id , title: item.productName,image: item.imageUrl , price: item.price, rating: (item.rating /20).roundToDouble(),compparedItems:compparedItems , item : item),
                      );
                    }, childCount: _savedItems.length),
                  ),
                ),
                              ],
                            ),
                          ),
                          Container(
                            width: 40,
                            height: 4,
                            margin: const EdgeInsets.only(top: 4),
                            decoration: BoxDecoration(
                              color: Colors.grey[300],
                              borderRadius: BorderRadius.circular(2),
                            ),
                          ),
                        ],
                      );
                    },
                  ),
                );
              },
            );
      },
      child: Container(
        padding: EdgeInsets.symmetric(horizontal: 16.0, vertical: 12.0),
        margin: EdgeInsets.symmetric(horizontal: 16.0),
        decoration: BoxDecoration(
          gradient: LinearGradient(
            colors: [Colors.white, Colors.grey.shade300],
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
        child: Row(
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
                  cut(widget.text),
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
