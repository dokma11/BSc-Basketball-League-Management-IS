import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TradeManagementComponent } from './trade-management.component';

describe('TradeManagementComponent', () => {
  let component: TradeManagementComponent;
  let fixture: ComponentFixture<TradeManagementComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TradeManagementComponent]
    });
    fixture = TestBed.createComponent(TradeManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
