import 'package:flutter/material.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';

import 'package:shopallymobile/features/comparing/domain/Entity/comparison_entity.dart';

class ComparisonTableWidget extends StatelessWidget {
  final List<ComparisonEntity> comparisonEntity;
  late final List<String> features;
  late final List<Map<String, dynamic>> products;

  ComparisonTableWidget({super.key, required this.comparisonEntity}) {
    // Initialize in constructor
    features = comparisonEntity.isNotEmpty
        ? (comparisonEntity[0].table.features)
        : [];

    if (features.isNotEmpty) {
      features.remove('title');
    }

    products = comparisonEntity.map((e) {
      final productMap = Map<String, dynamic>.from(
        e.table.featureValuePair ?? {},
      );
      productMap['name'] = e.product.title;
      return productMap;
    }).toList();
  }

  TextStyle _getFeatureStyle(String feature) {
    if (feature == 'Price') {
      return const TextStyle(
        color: Colors.amber,
        fontWeight: FontWeight.w700,
        fontSize: 16,
      );
    } else if (feature == 'Rating') {
      return const TextStyle(
        color: Colors.black,
        fontSize: 16,
        fontWeight: FontWeight.w700,
      );
    } else {
      return const TextStyle(
        color: Color.fromRGBO(117, 123, 129, 1),
        fontWeight: FontWeight.w400,
        fontSize: 14,
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(15.0),
      child: Container(
        decoration: const BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.all(Radius.circular(10)),
          boxShadow: [
            BoxShadow(
              color: Color.fromRGBO(0, 0, 0, 0.5),
              spreadRadius: 1,
              blurRadius: 10,
            ),
          ],
        ),
        child: Scrollbar(
          thickness: 8,
          thumbVisibility: true,

          child: SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisSize: MainAxisSize.min,
              children: [
                Container(
                  margin: const EdgeInsets.all(30),
                  child:  Text(
                    getText('Detailed Comparison'),
                    style: TextStyle(
                      color: Color.fromRGBO(38, 43, 50, 1),
                      fontWeight: FontWeight.w600,
                      fontSize: 20,
                    ),
                  ),
                ),
                Container(
                  margin: const EdgeInsets.fromLTRB(30, 0, 30, 0),
                  child: DataTable(
                    columns: [
                      const DataColumn(
                        label: Text(
                          "Feature",
                          style: TextStyle(
                            fontSize: 16,
                            fontWeight: FontWeight.w500,
                            color: Color.fromRGBO(38, 43, 50, 1),
                          ),
                        ),
                      ),
                      for (var product in products)
                        DataColumn(
                          label: SizedBox(
                            width: 200,
                            child: Text(
                              product['name'],
                              style: const TextStyle(
                                fontSize: 16,
                                fontWeight: FontWeight.w500,
                                color: Color.fromRGBO(38, 43, 50, 1),
                              ),
                              maxLines: 1,
                              overflow: TextOverflow.ellipsis,
                            ),
                          ),
                        ),
                    ],
                    rows: [
                      for (var feature in features)
                        DataRow(
                          cells: [
                            DataCell(
                              SizedBox(
                                width: 150,
                                child: Text(
                                  feature,
                                  style: const TextStyle(
                                    fontSize: 16,
                                    fontWeight: FontWeight.w500,
                                    color: Color.fromRGBO(38, 43, 50, 1),
                                  ),
                                  softWrap: true,
                                ),
                              ),
                            ),
                            for (var product in products)
                              DataCell(
                                Center(
                                  child: SizedBox(
                                    width: 200,

                                    child: Text(
                                      product[feature] ?? "",
                                      style: _getFeatureStyle(feature),
                                    ),
                                  ),
                                ),
                              ),
                          ],
                        ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
