import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TradeRequestCardComponent } from './trade-request-card.component';

describe('TradeRequestCardComponent', () => {
  let component: TradeRequestCardComponent;
  let fixture: ComponentFixture<TradeRequestCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TradeRequestCardComponent]
    });
    fixture = TestBed.createComponent(TradeRequestCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
