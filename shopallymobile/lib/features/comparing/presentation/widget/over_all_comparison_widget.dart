import 'package:flutter/material.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/over_all_comparison_entity.dart';

class OverAllComparisonWidget extends StatelessWidget {
  final OverAllComparisonEntity overAllComparisonEntity;
  const OverAllComparisonWidget({
    super.key,
    required this.overAllComparisonEntity,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {},
      child: Padding(
        padding: const EdgeInsets.all(15.0),
        child: Column(
          children: [
            Row(
              spacing: 15,
              children: [
                Icon(Icons.star, size: 24),
                Text(
                  'AI Recommendation',
                  style: TextStyle(fontWeight: FontWeight.w600, fontSize: 20),
                ),
              ],
            ),
            SizedBox(height: 15,),
            Container(
              decoration: BoxDecoration(
                color: Colors.grey.shade100,
                borderRadius: BorderRadius.all(Radius.circular(10)),
              ),
              child: Padding(
                padding: const EdgeInsets.all(15.0),
                child: Row(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  spacing: 20,
                  children: [
                    CircleAvatar(
                      backgroundColor: Color.fromRGBO(255, 211, 0, 1),
                      child: Icon(Icons.flash_on_sharp, size: 24),
                    ),
                    Expanded(
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        spacing: 10,
                        children: [
                          Text(
                            overAllComparisonEntity.bestValueProduct,
                            maxLines: 1,
                            overflow: TextOverflow.ellipsis,
                            style: TextStyle(
                              fontWeight: FontWeight.w600,
                              fontSize: 16,
                              color: Color.fromRGBO(38, 43, 50, 1),
                            ),
                          ),
                          SizedBox(height: 10),
                          Text(
                            overAllComparisonEntity.summary,
                            softWrap: true,
                            style: TextStyle(
                              fontWeight: FontWeight.w400,
                              fontSize: 14,
                              color: Color.fromRGBO(117, 123, 129, 1),
                            ),
                          ),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.start,
                            spacing: 10,
                            children: [
                              Text(
                                '${overAllComparisonEntity.bestValuePrice.etb} ETB',
                                style: TextStyle(
                                  fontSize: 14,
                                  fontWeight: FontWeight.w500,
                                  color: Color.fromRGBO(255, 211, 0, 1),
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
          ],
        ),
      ),
    );
  }
}
