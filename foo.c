// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mutex;

// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&mutex);
        i++;
        pthread_mutex_unlock(&mutex);
    }
    printf("this is i incremented: %d\n", i);
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&mutex);
        i--;
        pthread_mutex_unlock(&mutex);
        }
    printf("this is i decremented: %d\n", i);
    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    
    pthread_t thread_1, thread_2;
    pthread_mutex_init(&mutex, NULL);

    if (pthread_create(&thread_1, NULL, incrementingThreadFunction, NULL) != 0){
        fprintf(stderr,"error has happended with incrementing \n");
        return 1;
    }

    if (pthread_create(&thread_2, NULL, decrementingThreadFunction, NULL) != 0) {
        fprintf(stderr, "Error creating thread2\n");
        return 1;
    }

    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`  

    pthread_join(thread_1, NULL);
    pthread_join(thread_2, NULL);  

    printf("The magic number is: %d\n", i);

    pthread_mutex_destroy(&mutex);
    return 0;
}

