 #include <stdio.h>
#include <stdlib.h>

 /* Definition for an interval.
 */
  struct Interval {
      int start;
      int end;
  };
 
 typedef struct Interval interval;

 interval* insert(interval *intervals, int sz, interval newInterval, int *len);

void main() {
	interval ints1[2] = {{1,3}, {4,6}};
	interval new1 = {2, 5};

	interval ints2[2] = {{1,2}, {4,6}};
	interval new2 = {3, 5};

	interval ints3[2] = {{1,3}, {5,6}};
	interval new3 = {2, 4};
	
	interval ints4[2] = {{1,2}, {5,6}};
	interval new4 = {3, 4};
	
	interval ints5[2] = {{3,4}, {5,6}};
	interval new5 = {1,2};
	
	interval ints6[2] = {{3,4}, {5,6}};
	interval new6 = {1,3};
	
	interval ints7[2] = {{3,4}, {5,6}};
	interval new7 = {1,5};
	
	interval ints8[1] = {{1,2}};
	interval new8 = {3,4};
	
	interval ints9[1] = {{1,4}};
	interval new9 = {2,5};

	interval ints10[2] = {{1,2}, {3,4}};
	interval new10 = {2,5};

	interval ints11[2] = {{1,2}, {5,6}};
	interval new11 = {3,7};
	
	interval ints12[2] = {{3,4}, {6,7}};
	interval new12 = {1,5};
	
	interval ints13[2] = {{3,4}, {5,6}};
	interval new13 = {1,7};
	
	interval ints14[1] = {{1,4}};
	interval new14 = {2,3};

	interval* result;
	int len;

   	printf("\n ints1:\n");
	result = insert(&ints1[0], 2, new1, &len);
   	printf("\n ints2:\n");
	result = insert(&ints2[0], 2, new2, &len);
   	printf("\n ints3:\n");
	result = insert(&ints3[0], 2, new3, &len);
   	printf("\n ints4:\n");
	result = insert(&ints4[0], 2, new4, &len);
   	printf("\n ints5:\n");
	result = insert(&ints5[0], 2, new5, &len);
   	printf("\n ints6:\n");
	result = insert(&ints6[0], 2, new6, &len);
   	printf("\n ints7:\n");
	result = insert(&ints7[0], 2, new7, &len);
   	printf("\n ints8:\n");
	result = insert(&ints8[0], 1, new8, &len);
   	printf("\n ints9:\n");
	result = insert(&ints9[0], 1, new9, &len);
   	printf("\n ints10:\n");
	result = insert(&ints10[0], 2, new10, &len);
   	printf("\n ints11:\n");
	result = insert(&ints11[0], 2, new11, &len);
   	printf("\n ints12:\n");
	result = insert(&ints12[0], 2, new12, &len);
   	printf("\n ints13:\n");
	result = insert(&ints13[0], 2, new13, &len);
   	printf("\n ints14:\n");
	result = insert(&ints14[0], 1, new14, &len);
	// result = insert(&ints1[0], 2, new1, &len);

}
/*
 * intervals : the array of interval
 * sz : number of entries in intervals
 * newInterval : new Interval to be inserted
 * len : populate the length of returned array of intervals in len
 */
interval* insert(interval *intervals, int sz, interval newInterval, int *len) {
	int item;
	interval *res = (interval*) malloc((sz+1) * sizeof(interval));

	int added = 0, count = 0;
	for (int i = 0; i < sz; i++) {
		interval cur = *(intervals + i);
		if (cur.end < newInterval.start) {
			res[count] = cur;
			count++;
			continue;
		}
		if (cur.start < newInterval.start) {
			newInterval.start = cur.start;
		}

		if (cur.start > newInterval.end) {
			res[count] = newInterval;
			count++;
			item = i;
			added = 1;
			break;
		}

		if (cur.end >= newInterval.end) {
			res[count].start = newInterval.start;
			res[count].end = cur.end;
			count++;
			item = i + 1;
			added = 1;
			break;
		}
	}

	if (added == 0) {
		res[count] = newInterval;
		count++;	
	}else{
		for (int i = item; i < sz; i++){
			interval cur = *(intervals + i);
			res[count] = cur;
			count++;
		}
	}

	*len = count;

   printf("len: %d\n", *len);
   printf("result:\n");
   for (int i = 0; i < count; i++){
   	printf("[%d %d] ", (*(res+i)).start, (*(res+i)).end);
   }
   printf("\n");

	return res;
}
