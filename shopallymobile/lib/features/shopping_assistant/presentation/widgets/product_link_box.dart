import 'package:flutter/material.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';
import 'package:shopallymobile/features/saveditem/presentation/pages/detail_page.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/productlist/buttom_nav.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/productlist/product_info.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/features/comparing/presentation/bloc/compare_bloc.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart'
    as cmp;
import 'package:shopallymobile/features/comparing/domain/Entity/price_entity.dart'
    as cmp;

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
                        decoration: BoxDecoration(
                          color: Theme.of(context).cardColor,
                          borderRadius: const BorderRadius.only(
                            topLeft: Radius.circular(20),
                            topRight: Radius.circular(20),
                          ),
                        ),
                        child: CustomScrollView(
                          controller: scrollController,
                          slivers: <Widget>[
                            SliverAppBar(
                              pinned: true,
                              backgroundColor: Theme.of(context).cardColor,
                              automaticallyImplyLeading: false,
                              elevation: 0,
                              title: Row(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceBetween,
                                children: [
                                  IconButton(
                                    icon: Icon(
                                      Icons.arrow_back_ios,
                                      color: Theme.of(
                                        context,
                                      ).textTheme.bodyLarge?.color,
                                    ),
                                    onPressed: () {
                                      Navigator.of(context).pop();
                                    },
                                  ),
                                  Text(
                                    getText('Product Search'),
                                    style: TextStyle(
                                      fontSize: 18,
                                      fontWeight: FontWeight.bold,
                                      color: Theme.of(
                                        context,
                                      ).textTheme.bodyLarge?.color,
                                    ),
                                  ),
                                  TextButton(
                                    onPressed: () {
                                      // Validate selection count
                                      if (compparedItems.length < 2) {
                                        ScaffoldMessenger.of(
                                          context,
                                        ).showSnackBar(
                                          SnackBar(
                                            content: Text(
                                              getText(
                                                'Select at least 2 items to compare.',
                                              ),
                                            ),
                                            duration: const Duration(
                                              seconds: 2,
                                            ),
                                          ),
                                        );
                                        return;
                                      }

                                      // Map shopping assistant ProductEntity -> comparing ProductEntity
                                      final List<cmp.ProductEntity>
                                      mapped = compparedItems.map((item) {
                                        final fxTs =
                                            item.fxTimestamp
                                                ?.toIso8601String() ??
                                            DateTime.now().toIso8601String();
                                        final price = cmp.PriceEntity(
                                          item.priceEtb,
                                          item.priceUsd,
                                          fxTs,
                                        );
                                        return cmp.ProductEntity(
                                          imageUrl: item.imageUrl,
                                          price: price,
                                          id: item.id,
                                          title: item.productName,
                                          aiMatchPercentage: item
                                              .aiMatchPercentage
                                              .round(),
                                          productRating: item.productRating,
                                          deepLinkUrl: item.deeplinkUrl,
                                          deliveryEstimate:
                                              item.deliveryEstimate,
                                          description: item.productDescription,
                                          productSmallImageUrls: null,
                                          numberSold: item.numberSold,
                                          summaryBullets: item.summaryBullets,
                                          tax: item.taxRate,
                                          discount: item.discount,
                                        );
                                      }).toList();

                                      // Dispatch compare event
                                      context.read<CompareBloc>().add(
                                        CompareProductsEvent(mapped),
                                      );

                                      // Navigate to compare screen which listens to results
                                      Navigator.pushNamed(
                                        context,
                                        '/compare-products',
                                      );
                                    },
                                    child: Text(
                                      getText('Compare'),
                                      style: TextStyle(
                                        color: Theme.of(
                                          context,
                                        ).textTheme.bodyLarge?.color,
                                      ),
                                    ),
                                  ),
                                ],
                              ),
                              bottom: PreferredSize(
                                preferredSize: const Size.fromHeight(40),
                                child: Align(
                                  alignment: Alignment.centerLeft,
                                  child: Padding(
                                    padding: const EdgeInsets.only(
                                      left: 16.0,
                                      bottom: 8.0,
                                    ),
                                    child: Text(
                                      '',
                                      style: Theme.of(context)
                                          .textTheme
                                          .titleLarge
                                          ?.copyWith(
                                            fontWeight: FontWeight.bold,
                                            color: Theme.of(
                                              context,
                                            ).textTheme.bodyLarge?.color,
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
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: [
                                    SizedBox(height: 8),
                                    Text(getText('Search results for')),
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
                                delegate: SliverChildBuilderDelegate((
                                  context,
                                  index,
                                ) {
                                  final item = _savedItems[index];
                                  return GestureDetector(
                                    onTap: () {
                                      Navigator.push(
                                        context,
                                        MaterialPageRoute(
                                          builder: (context) =>
                                              DetailPage(item: item),
                                        ),
                                      );
                                    },

                                    child: ProductInfo(
                                      id: item.id,
                                      title: item.productName,
                                      image: item.imageUrl,
                                      price: item.price,
                                      rating: (item.rating / 20)
                                          .roundToDouble(),
                                      compparedItems: compparedItems,
                                      item: item,
                                    ),
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
                          color: Theme.of(context).dividerColor,
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
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 12.0),
        margin: const EdgeInsets.symmetric(horizontal: 16.0),
        decoration: BoxDecoration(
          color: Theme.of(context).cardColor,
          borderRadius: const BorderRadius.all(Radius.circular(12.0)),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withOpacity(
                Theme.of(context).brightness == Brightness.dark ? 0.2 : 0.06,
              ),
              blurRadius: 8.0,
              offset: const Offset(0, 2),
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
                    Icon(
                      Icons.link,
                      color: Theme.of(context).colorScheme.primary,
                    ),
                    const SizedBox(width: 8.0),
                    Text(
                      getText('Product Search'),
                      style: TextStyle(
                        fontSize: 14.0,
                        fontWeight: FontWeight.w400,
                        color: Theme.of(
                          context,
                        ).textTheme.bodyLarge?.color?.withOpacity(0.8),
                        fontFamily: 'Arial',
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 4.0),
                Text(
                  cut(widget.text),
                  style: TextStyle(
                    fontSize: 14.0,
                    color: Theme.of(context).textTheme.bodyLarge?.color,
                    fontWeight: FontWeight.w500,
                    fontFamily: 'Arial',
                  ),
                ),
              ],
            ),
            Icon(
              Icons.arrow_forward_ios,
              color: Theme.of(context).textTheme.bodyLarge?.color,
              size: 16.0,
            ),
          ],
        ),
      ),
    );
  }
}
