#include <pthread.h>
#include <stdio.h>
#include <mutex>

std::mutex mtx;

int i = 0;


pthread_mutex_t lock;
/**
 * Mutex is thread-otiented and semaphores are token oriented. Since we in this case is
 * working on two particular threads it seams reasonable to use mutex. Sem_t might be
 * prefered when we use interrupt handlers.
 */
void *incrementingThreadFunction(void* num) {
    for (int j = 0; j < 1000099; j++) {
        mtx.lock();                // Acquire lock
        (*(int*)num)++;              // Critical work
        mtx.unlock();           // Release lock
    }

    return NULL;
}

void *decrementingThreadFunction(void* num) {
    for (int j = 0; j < 1000000; j++) {
        mtx.lock();;
        (*(int*)num)--;
        mtx.unlock();
    }

    return NULL;
}


int main() {
    pthread_mutex_init(&lock, NULL);

    pthread_t incrementingThread, decrementingThread;

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, &i);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, &i);

    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);

    printf("The magic number is: %d\n", i);

    return 0;
