using namespace std;
#include <cstdlib>
#include <iostream>
int partition(int array[], int start, int end);
void quicksort(int array[], int start, int end);
int main(void) {
  int arr[100];
  srand(1230091);
  for (int i = 0; i < 100; i++) {
    arr[i] = rand() % 100 + 1;
  }

  for (int i = 0; i < 100; i++) {
    cout << arr[i] << "\n";
  }
  quicksort(arr, 0, 99);
  for (int i = 0; i < 100; i++) {
    cout << arr[i] << "\n";
  }
  return 1;
}

int partition(int array[], int start, int end) {
  int i = start - 1;
  for (int j = start; j < end; j++) {
    if (array[j] < array[end]) {
      i++;
      int temp = array[i];
      array[i] = array[j];
      array[j] = temp;
    }
  }
  i++;
  int temp = array[i];
  array[i] = array[end];
  array[end] = temp;
  return i;
}
void quicksort(int array[], int start, int end) {
  if (end <= start)
    return;
  int pivot = partition(array, start, end);
  quicksort(array, start, pivot - 1);
  quicksort(array, pivot + 1, end);
}
