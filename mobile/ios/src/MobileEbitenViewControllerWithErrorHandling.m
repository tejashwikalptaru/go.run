//
//  MobileEbitenViewWithErrorHandling.m
//  adventure runner
//
//  Created by Tejashwi Kalp Taru
//

#import "MobileEbitenViewControllerWithErrorHandling.h"

#import <Foundation/Foundation.h>

@implementation MobileEbitenViewControllerWithErrorHandling {
}

- (void)onErrorOnGameUpdate:(NSError*)err {
    // You can define your own error handling e.g., using Crashlytics.
    NSLog(@"Inovation Error!: %@", err);
}

@end
